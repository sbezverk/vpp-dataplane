// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package cnat

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/vpe"
)

// RPCService defines RPC service cnat.
type RPCService interface {
	CnatAddDelSnatPrefix(ctx context.Context, in *CnatAddDelSnatPrefix) (*CnatAddDelSnatPrefixReply, error)
	CnatGetSnatAddresses(ctx context.Context, in *CnatGetSnatAddresses) (*CnatGetSnatAddressesReply, error)
	CnatSessionDump(ctx context.Context, in *CnatSessionDump) (RPCService_CnatSessionDumpClient, error)
	CnatSessionPurge(ctx context.Context, in *CnatSessionPurge) (*CnatSessionPurgeReply, error)
	CnatSetSnatAddresses(ctx context.Context, in *CnatSetSnatAddresses) (*CnatSetSnatAddressesReply, error)
	CnatTranslationDel(ctx context.Context, in *CnatTranslationDel) (*CnatTranslationDelReply, error)
	CnatTranslationDump(ctx context.Context, in *CnatTranslationDump) (RPCService_CnatTranslationDumpClient, error)
	CnatTranslationUpdate(ctx context.Context, in *CnatTranslationUpdate) (*CnatTranslationUpdateReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) CnatAddDelSnatPrefix(ctx context.Context, in *CnatAddDelSnatPrefix) (*CnatAddDelSnatPrefixReply, error) {
	out := new(CnatAddDelSnatPrefixReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CnatGetSnatAddresses(ctx context.Context, in *CnatGetSnatAddresses) (*CnatGetSnatAddressesReply, error) {
	out := new(CnatGetSnatAddressesReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CnatSessionDump(ctx context.Context, in *CnatSessionDump) (RPCService_CnatSessionDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_CnatSessionDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_CnatSessionDumpClient interface {
	Recv() (*CnatSessionDetails, error)
	api.Stream
}

type serviceClient_CnatSessionDumpClient struct {
	api.Stream
}

func (c *serviceClient_CnatSessionDumpClient) Recv() (*CnatSessionDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *CnatSessionDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) CnatSessionPurge(ctx context.Context, in *CnatSessionPurge) (*CnatSessionPurgeReply, error) {
	out := new(CnatSessionPurgeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CnatSetSnatAddresses(ctx context.Context, in *CnatSetSnatAddresses) (*CnatSetSnatAddressesReply, error) {
	out := new(CnatSetSnatAddressesReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CnatTranslationDel(ctx context.Context, in *CnatTranslationDel) (*CnatTranslationDelReply, error) {
	out := new(CnatTranslationDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CnatTranslationDump(ctx context.Context, in *CnatTranslationDump) (RPCService_CnatTranslationDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_CnatTranslationDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_CnatTranslationDumpClient interface {
	Recv() (*CnatTranslationDetails, error)
	api.Stream
}

type serviceClient_CnatTranslationDumpClient struct {
	api.Stream
}

func (c *serviceClient_CnatTranslationDumpClient) Recv() (*CnatTranslationDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *CnatTranslationDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) CnatTranslationUpdate(ctx context.Context, in *CnatTranslationUpdate) (*CnatTranslationUpdateReply, error) {
	out := new(CnatTranslationUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
