// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package define

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type SaleService interface {
	// Parameters:
	//  - ID
	GetSubOrder(id int64) (r *SubOrder, err error)
	// Parameters:
	//  - OrderNo
	GetSubOrderByNo(orderNo string) (r *SubOrder, err error)
	// Parameters:
	//  - SubOrderId
	GetSubOrderItems(subOrderId int64) (r []*OrderItem, err error)
}

type SaleServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewSaleServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SaleServiceClient {
	return &SaleServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewSaleServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SaleServiceClient {
	return &SaleServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - ID
func (p *SaleServiceClient) GetSubOrder(id int64) (r *SubOrder, err error) {
	if err = p.sendGetSubOrder(id); err != nil {
		return
	}
	return p.recvGetSubOrder()
}

func (p *SaleServiceClient) sendGetSubOrder(id int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetSubOrder", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SaleServiceGetSubOrderArgs{
		ID: id,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SaleServiceClient) recvGetSubOrder() (value *SubOrder, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetSubOrder" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetSubOrder failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetSubOrder failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error150 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error151 error
		error151, err = error150.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error151
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetSubOrder failed: invalid message type")
		return
	}
	result := SaleServiceGetSubOrderResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - OrderNo
func (p *SaleServiceClient) GetSubOrderByNo(orderNo string) (r *SubOrder, err error) {
	if err = p.sendGetSubOrderByNo(orderNo); err != nil {
		return
	}
	return p.recvGetSubOrderByNo()
}

func (p *SaleServiceClient) sendGetSubOrderByNo(orderNo string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetSubOrderByNo", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SaleServiceGetSubOrderByNoArgs{
		OrderNo: orderNo,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SaleServiceClient) recvGetSubOrderByNo() (value *SubOrder, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetSubOrderByNo" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetSubOrderByNo failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetSubOrderByNo failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error152 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error153 error
		error153, err = error152.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error153
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetSubOrderByNo failed: invalid message type")
		return
	}
	result := SaleServiceGetSubOrderByNoResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - SubOrderId
func (p *SaleServiceClient) GetSubOrderItems(subOrderId int64) (r []*OrderItem, err error) {
	if err = p.sendGetSubOrderItems(subOrderId); err != nil {
		return
	}
	return p.recvGetSubOrderItems()
}

func (p *SaleServiceClient) sendGetSubOrderItems(subOrderId int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetSubOrderItems", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SaleServiceGetSubOrderItemsArgs{
		SubOrderId: subOrderId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SaleServiceClient) recvGetSubOrderItems() (value []*OrderItem, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetSubOrderItems" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetSubOrderItems failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetSubOrderItems failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error154 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error155 error
		error155, err = error154.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error155
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetSubOrderItems failed: invalid message type")
		return
	}
	result := SaleServiceGetSubOrderItemsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type SaleServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SaleService
}

func (p *SaleServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SaleServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SaleServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSaleServiceProcessor(handler SaleService) *SaleServiceProcessor {

	self156 := &SaleServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self156.processorMap["GetSubOrder"] = &saleServiceProcessorGetSubOrder{handler: handler}
	self156.processorMap["GetSubOrderByNo"] = &saleServiceProcessorGetSubOrderByNo{handler: handler}
	self156.processorMap["GetSubOrderItems"] = &saleServiceProcessorGetSubOrderItems{handler: handler}
	return self156
}

func (p *SaleServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x157 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x157.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x157

}

type saleServiceProcessorGetSubOrder struct {
	handler SaleService
}

func (p *saleServiceProcessorGetSubOrder) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SaleServiceGetSubOrderArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetSubOrder", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SaleServiceGetSubOrderResult{}
	var retval *SubOrder
	var err2 error
	if retval, err2 = p.handler.GetSubOrder(args.ID); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetSubOrder: "+err2.Error())
		oprot.WriteMessageBegin("GetSubOrder", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetSubOrder", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type saleServiceProcessorGetSubOrderByNo struct {
	handler SaleService
}

func (p *saleServiceProcessorGetSubOrderByNo) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SaleServiceGetSubOrderByNoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetSubOrderByNo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SaleServiceGetSubOrderByNoResult{}
	var retval *SubOrder
	var err2 error
	if retval, err2 = p.handler.GetSubOrderByNo(args.OrderNo); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetSubOrderByNo: "+err2.Error())
		oprot.WriteMessageBegin("GetSubOrderByNo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetSubOrderByNo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type saleServiceProcessorGetSubOrderItems struct {
	handler SaleService
}

func (p *saleServiceProcessorGetSubOrderItems) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SaleServiceGetSubOrderItemsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetSubOrderItems", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SaleServiceGetSubOrderItemsResult{}
	var retval []*OrderItem
	var err2 error
	if retval, err2 = p.handler.GetSubOrderItems(args.SubOrderId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetSubOrderItems: "+err2.Error())
		oprot.WriteMessageBegin("GetSubOrderItems", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetSubOrderItems", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - ID
type SaleServiceGetSubOrderArgs struct {
	ID int64 `thrift:"id,1" json:"id"`
}

func NewSaleServiceGetSubOrderArgs() *SaleServiceGetSubOrderArgs {
	return &SaleServiceGetSubOrderArgs{}
}

func (p *SaleServiceGetSubOrderArgs) GetID() int64 {
	return p.ID
}
func (p *SaleServiceGetSubOrderArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *SaleServiceGetSubOrderArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSubOrder_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *SaleServiceGetSubOrderArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaleServiceGetSubOrderArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SaleServiceGetSubOrderResult struct {
	Success *SubOrder `thrift:"success,0" json:"success,omitempty"`
}

func NewSaleServiceGetSubOrderResult() *SaleServiceGetSubOrderResult {
	return &SaleServiceGetSubOrderResult{}
}

var SaleServiceGetSubOrderResult_Success_DEFAULT *SubOrder

func (p *SaleServiceGetSubOrderResult) GetSuccess() *SubOrder {
	if !p.IsSetSuccess() {
		return SaleServiceGetSubOrderResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SaleServiceGetSubOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SaleServiceGetSubOrderResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &SubOrder{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSubOrder_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SaleServiceGetSubOrderResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaleServiceGetSubOrderResult(%+v)", *p)
}

// Attributes:
//  - OrderNo
type SaleServiceGetSubOrderByNoArgs struct {
	OrderNo string `thrift:"orderNo,1" json:"orderNo"`
}

func NewSaleServiceGetSubOrderByNoArgs() *SaleServiceGetSubOrderByNoArgs {
	return &SaleServiceGetSubOrderByNoArgs{}
}

func (p *SaleServiceGetSubOrderByNoArgs) GetOrderNo() string {
	return p.OrderNo
}
func (p *SaleServiceGetSubOrderByNoArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderByNoArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.OrderNo = v
	}
	return nil
}

func (p *SaleServiceGetSubOrderByNoArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSubOrderByNo_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderByNoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("orderNo", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:orderNo: ", p), err)
	}
	if err := oprot.WriteString(string(p.OrderNo)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.orderNo (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:orderNo: ", p), err)
	}
	return err
}

func (p *SaleServiceGetSubOrderByNoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaleServiceGetSubOrderByNoArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SaleServiceGetSubOrderByNoResult struct {
	Success *SubOrder `thrift:"success,0" json:"success,omitempty"`
}

func NewSaleServiceGetSubOrderByNoResult() *SaleServiceGetSubOrderByNoResult {
	return &SaleServiceGetSubOrderByNoResult{}
}

var SaleServiceGetSubOrderByNoResult_Success_DEFAULT *SubOrder

func (p *SaleServiceGetSubOrderByNoResult) GetSuccess() *SubOrder {
	if !p.IsSetSuccess() {
		return SaleServiceGetSubOrderByNoResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SaleServiceGetSubOrderByNoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SaleServiceGetSubOrderByNoResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderByNoResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &SubOrder{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderByNoResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSubOrderByNo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderByNoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SaleServiceGetSubOrderByNoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaleServiceGetSubOrderByNoResult(%+v)", *p)
}

// Attributes:
//  - SubOrderId
type SaleServiceGetSubOrderItemsArgs struct {
	SubOrderId int64 `thrift:"subOrderId,1" json:"subOrderId"`
}

func NewSaleServiceGetSubOrderItemsArgs() *SaleServiceGetSubOrderItemsArgs {
	return &SaleServiceGetSubOrderItemsArgs{}
}

func (p *SaleServiceGetSubOrderItemsArgs) GetSubOrderId() int64 {
	return p.SubOrderId
}
func (p *SaleServiceGetSubOrderItemsArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderItemsArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.SubOrderId = v
	}
	return nil
}

func (p *SaleServiceGetSubOrderItemsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSubOrderItems_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderItemsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("subOrderId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:subOrderId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.SubOrderId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.subOrderId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:subOrderId: ", p), err)
	}
	return err
}

func (p *SaleServiceGetSubOrderItemsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaleServiceGetSubOrderItemsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SaleServiceGetSubOrderItemsResult struct {
	Success []*OrderItem `thrift:"success,0" json:"success,omitempty"`
}

func NewSaleServiceGetSubOrderItemsResult() *SaleServiceGetSubOrderItemsResult {
	return &SaleServiceGetSubOrderItemsResult{}
}

var SaleServiceGetSubOrderItemsResult_Success_DEFAULT []*OrderItem

func (p *SaleServiceGetSubOrderItemsResult) GetSuccess() []*OrderItem {
	return p.Success
}
func (p *SaleServiceGetSubOrderItemsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SaleServiceGetSubOrderItemsResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderItemsResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*OrderItem, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem158 := &OrderItem{}
		if err := _elem158.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem158), err)
		}
		p.Success = append(p.Success, _elem158)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderItemsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSubOrderItems_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SaleServiceGetSubOrderItemsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SaleServiceGetSubOrderItemsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaleServiceGetSubOrderItemsResult(%+v)", *p)
}
