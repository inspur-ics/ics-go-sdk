package icsgo

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client"
    "github.com/inspur-ics/ics-go-sdk/client/restful"
    "github.com/inspur-ics/ics-go-sdk/session"
    "k8s.io/klog"
    "net"
    "net/url"
    "sync"
)

type ICSConnection struct {
    Client            *client.Client
    Username          string
    Password          string
    Hostname          string
    Port              string
    Insecure          bool
    credentialsLock   sync.Mutex
}

var (
    clientLock sync.Mutex
)

func (connection *ICSConnection) Connect(ctx context.Context) error {
    var err error
    clientLock.Lock()
    defer clientLock.Unlock()

    if connection.Client == nil {
        connection.Client, err = connection.NewClient(ctx)
        if err != nil {
            klog.Errorf("Failed to create ics-go-sdk client. err: %+v", err)
            return err
        }
        return nil
    }
    m := session.NewManager(connection.Client)
    userSession, err := m.UserSession(ctx)
    if err != nil {
        klog.Errorf("Error while obtaining user session. err: %+v", err)
        return err
    }
    if userSession != nil {
        return nil
    }
    klog.Warning("Creating new client session since the existing session is not valid or not authenticated")

    connection.Client, err = connection.NewClient(ctx)
    if err != nil {
        klog.Errorf("Failed to create ics-go-sdk client. err: %+v", err)
        return err
    }
    return nil
}


// login calls SessionManager.LoginByToken if certificate and private key are configured,
// otherwise calls SessionManager.Login with user and password.
func (connection *ICSConnection) login(ctx context.Context, client *client.Client) error {
    m := session.NewManager(client)
    connection.credentialsLock.Lock()
    defer connection.credentialsLock.Unlock()

    klog.V(3).Infof("SessionManager.Login with username %q", connection.Username)
    return m.Login(ctx, url.UserPassword(connection.Username, connection.Password))
}

// NewClient creates a new ics-go-sdk client for the ICSConnection obj
func (connection *ICSConnection) NewClient(ctx context.Context) (*client.Client, error) {
    url, err := restful.ParseURL(net.JoinHostPort(connection.Hostname, connection.Port))
    if err != nil {
        klog.Errorf("Failed to parse URL: %s. err: %+v", url, err)
        return nil, err
    }

    sc := restful.NewClient(url, connection.Insecure)

    client, err := client.NewClient(ctx, sc)
    if err != nil {
        klog.Errorf("Failed to create new client. err: %+v", err)
        return nil, err
    }
    err = connection.login(ctx, client)
    if err != nil {
        return nil, err
    }
    return client, nil
}

// NewClient creates a new ics-go-sdk client for the ICSConnection obj
func (connection *ICSConnection) GetClient() (*client.Client, error) {
    var client = connection.Client
    var err error
    ctx := context.Background()
    if connection.Client == nil {
        err = connection.Connect(ctx)
        if err != nil {
            klog.Errorf("Failed to create ics-go-sdk client. err: %+v", err)
            return nil, err
        } else {
            client = connection.Client
        }
        return client, nil
    }
    m := session.NewManager(client)
    userSession, err := m.UserSession(ctx)
    if err != nil || userSession == nil {
        err = connection.Connect(ctx)
        if err != nil {
            klog.Errorf("Failed to create ics-go-sdk client. err: %+v", err)
            return nil, err
        } else {
            client = connection.Client
        }
        return client, nil
    }
    return client, nil
}