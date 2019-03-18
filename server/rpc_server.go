package server

type RpcTest struct {
	A, B int64
}

//RpcTest 返回两个数的之和
func (this RpcTest) Add(args RpcTest, r *int64) error {
	*r = args.A + args.B
	return nil
}

//Multply 返回两个数的之积
func (this RpcTest) Multply(args RpcTest, r *int64) error {
	*r = args.A + args.B
	return nil
}
