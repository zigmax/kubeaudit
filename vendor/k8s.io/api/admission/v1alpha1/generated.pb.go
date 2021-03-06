/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/api/admission/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/admission/v1alpha1/generated.proto

	It has these top-level messages:
		AdmissionReview
		AdmissionReviewSpec
		AdmissionReviewStatus
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import k8s_io_apiserver_pkg_admission "k8s.io/apiserver/pkg/admission"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *AdmissionReview) Reset()                    { *m = AdmissionReview{} }
func (*AdmissionReview) ProtoMessage()               {}
func (*AdmissionReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *AdmissionReviewSpec) Reset()                    { *m = AdmissionReviewSpec{} }
func (*AdmissionReviewSpec) ProtoMessage()               {}
func (*AdmissionReviewSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *AdmissionReviewStatus) Reset()                    { *m = AdmissionReviewStatus{} }
func (*AdmissionReviewStatus) ProtoMessage()               {}
func (*AdmissionReviewStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*AdmissionReview)(nil), "k8s.io.api.admission.v1alpha1.AdmissionReview")
	proto.RegisterType((*AdmissionReviewSpec)(nil), "k8s.io.api.admission.v1alpha1.AdmissionReviewSpec")
	proto.RegisterType((*AdmissionReviewStatus)(nil), "k8s.io.api.admission.v1alpha1.AdmissionReviewStatus")
}
func (m *AdmissionReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionReview) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n1, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n2, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *AdmissionReviewSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionReviewSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Kind.Size()))
	n3, err := m.Kind.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Object.Size()))
	n4, err := m.Object.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.OldObject.Size()))
	n5, err := m.OldObject.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Operation)))
	i += copy(dAtA[i:], m.Operation)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x32
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Namespace)))
	i += copy(dAtA[i:], m.Namespace)
	dAtA[i] = 0x3a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Resource.Size()))
	n6, err := m.Resource.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x42
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.SubResource)))
	i += copy(dAtA[i:], m.SubResource)
	dAtA[i] = 0x4a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.UserInfo.Size()))
	n7, err := m.UserInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	return i, nil
}

func (m *AdmissionReviewStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionReviewStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Allowed {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.Result != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Result.Size()))
		n8, err := m.Result.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AdmissionReview) Size() (n int) {
	var l int
	_ = l
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *AdmissionReviewSpec) Size() (n int) {
	var l int
	_ = l
	l = m.Kind.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Object.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.OldObject.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Operation)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Resource.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.SubResource)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.UserInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *AdmissionReviewStatus) Size() (n int) {
	var l int
	_ = l
	n += 2
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AdmissionReview) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionReview{`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "AdmissionReviewSpec", "AdmissionReviewSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "AdmissionReviewStatus", "AdmissionReviewStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AdmissionReviewSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionReviewSpec{`,
		`Kind:` + strings.Replace(strings.Replace(this.Kind.String(), "GroupVersionKind", "k8s_io_apimachinery_pkg_apis_meta_v1.GroupVersionKind", 1), `&`, ``, 1) + `,`,
		`Object:` + strings.Replace(strings.Replace(this.Object.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`OldObject:` + strings.Replace(strings.Replace(this.OldObject.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`Operation:` + fmt.Sprintf("%v", this.Operation) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Resource:` + strings.Replace(strings.Replace(this.Resource.String(), "GroupVersionResource", "k8s_io_apimachinery_pkg_apis_meta_v1.GroupVersionResource", 1), `&`, ``, 1) + `,`,
		`SubResource:` + fmt.Sprintf("%v", this.SubResource) + `,`,
		`UserInfo:` + strings.Replace(strings.Replace(this.UserInfo.String(), "UserInfo", "k8s_io_api_authentication_v1.UserInfo", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AdmissionReviewStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionReviewStatus{`,
		`Allowed:` + fmt.Sprintf("%v", this.Allowed) + `,`,
		`Result:` + strings.Replace(fmt.Sprintf("%v", this.Result), "Status", "k8s_io_apimachinery_pkg_apis_meta_v1.Status", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AdmissionReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdmissionReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdmissionReviewSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdmissionReviewSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionReviewSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Kind.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Object.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldObject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OldObject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operation = k8s_io_apiserver_pkg_admission.Operation(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubResource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubResource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdmissionReviewStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdmissionReviewStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionReviewStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Allowed = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &k8s_io_apimachinery_pkg_apis_meta_v1.Status{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/admission/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x4b, 0x1c, 0x49,
	0x14, 0xc7, 0xa7, 0x77, 0xc7, 0x71, 0xa6, 0x5c, 0xd6, 0xdd, 0x92, 0x85, 0x46, 0xd8, 0x56, 0x3c,
	0x2c, 0x2e, 0x68, 0x35, 0x1a, 0x23, 0x21, 0x90, 0x83, 0x03, 0x49, 0x08, 0x01, 0x0d, 0xa5, 0x86,
	0x90, 0x84, 0x40, 0x4d, 0xcf, 0x73, 0xa6, 0x32, 0xd3, 0x55, 0x4d, 0x55, 0xf5, 0x98, 0xdc, 0xf2,
	0x27, 0xe4, 0x90, 0x7f, 0x29, 0xe0, 0x25, 0xe0, 0xd1, 0x93, 0xc4, 0xc9, 0x7f, 0x91, 0x53, 0xe8,
	0xea, 0xea, 0x6e, 0x1d, 0x35, 0x89, 0x39, 0xcd, 0xbc, 0x1f, 0xdf, 0x4f, 0xbd, 0x7a, 0xf5, 0x5e,
	0xa3, 0x07, 0x83, 0x3b, 0x9a, 0x70, 0x19, 0x0e, 0xd2, 0x0e, 0x28, 0x01, 0x06, 0x74, 0x38, 0x02,
	0xd1, 0x95, 0x2a, 0x74, 0x01, 0x96, 0xf0, 0x90, 0x75, 0x63, 0xae, 0x35, 0x97, 0x22, 0x1c, 0xad,
	0xb1, 0x61, 0xd2, 0x67, 0x6b, 0x61, 0x0f, 0x04, 0x28, 0x66, 0xa0, 0x4b, 0x12, 0x25, 0x8d, 0xc4,
	0xff, 0xe6, 0xe9, 0x84, 0x25, 0x9c, 0x94, 0xe9, 0xa4, 0x48, 0x9f, 0x5f, 0xed, 0x71, 0xd3, 0x4f,
	0x3b, 0x24, 0x92, 0x71, 0xd8, 0x93, 0x3d, 0x19, 0x5a, 0x55, 0x27, 0x3d, 0xb0, 0x96, 0x35, 0xec,
	0xbf, 0x9c, 0x36, 0xbf, 0x72, 0xfe, 0xf0, 0xd4, 0xf4, 0x41, 0x18, 0x1e, 0x31, 0x93, 0x57, 0x30,
	0x79, 0xf6, 0xfc, 0x46, 0x95, 0x1d, 0xb3, 0xa8, 0xcf, 0x05, 0xa8, 0xb7, 0x61, 0x32, 0xe8, 0x65,
	0x0e, 0x1d, 0xc6, 0x60, 0xd8, 0x55, 0xaa, 0xf0, 0x3a, 0x95, 0x4a, 0x85, 0xe1, 0x31, 0x5c, 0x12,
	0x6c, 0xfe, 0x48, 0xa0, 0xa3, 0x3e, 0xc4, 0xec, 0x92, 0xee, 0xd6, 0x75, 0xba, 0xd4, 0xf0, 0x61,
	0xc8, 0x85, 0xd1, 0x46, 0x4d, 0x8a, 0x96, 0x3e, 0x7a, 0x68, 0x76, 0xab, 0xe8, 0x23, 0x85, 0x11,
	0x87, 0x43, 0xbc, 0x87, 0xea, 0x3a, 0x81, 0xc8, 0xf7, 0x16, 0xbd, 0xe5, 0x99, 0xf5, 0x75, 0xf2,
	0xdd, 0x96, 0x93, 0x09, 0xf5, 0x6e, 0x02, 0x51, 0xfb, 0x8f, 0xa3, 0xd3, 0x85, 0xda, 0xf8, 0x74,
	0xa1, 0x9e, 0x59, 0xd4, 0xd2, 0xf0, 0x4b, 0xd4, 0xd0, 0x86, 0x99, 0x54, 0xfb, 0xbf, 0x59, 0xee,
	0xc6, 0x0d, 0xb9, 0x56, 0xdb, 0xfe, 0xd3, 0x91, 0x1b, 0xb9, 0x4d, 0x1d, 0x73, 0xe9, 0xd3, 0x14,
	0x9a, 0xbb, 0xa2, 0x12, 0xfc, 0x0c, 0xd5, 0x07, 0x5c, 0x74, 0xdd, 0x5d, 0x36, 0xcf, 0x9d, 0x59,
	0xf6, 0x88, 0x24, 0x83, 0x5e, 0xe6, 0xd0, 0x24, 0x7b, 0x42, 0x32, 0x5a, 0x23, 0x0f, 0x95, 0x4c,
	0x93, 0xa7, 0xa0, 0x32, 0xd6, 0x63, 0x2e, 0xba, 0xd5, 0x7d, 0x32, 0x8b, 0x5a, 0x22, 0xde, 0x47,
	0x0d, 0xd9, 0x79, 0x0d, 0x91, 0x71, 0xf7, 0x59, 0xbd, 0x96, 0xed, 0xde, 0x8d, 0x50, 0x76, 0x78,
	0xff, 0x8d, 0x01, 0x91, 0x61, 0xab, 0x8b, 0xec, 0x58, 0x08, 0x75, 0x30, 0xfc, 0x0a, 0xb5, 0xe4,
	0xb0, 0x9b, 0x3b, 0xfd, 0xdf, 0x7f, 0x85, 0xfc, 0xb7, 0x23, 0xb7, 0x76, 0x0a, 0x0e, 0xad, 0x90,
	0xf8, 0x05, 0x6a, 0xc9, 0x24, 0x1b, 0x01, 0x2e, 0x85, 0x5f, 0x5f, 0xf4, 0x96, 0x5b, 0xed, 0x7b,
	0xa5, 0xa0, 0x08, 0x7c, 0x3d, 0x5d, 0x58, 0xae, 0xa6, 0x49, 0x83, 0x1a, 0x81, 0xca, 0x27, 0xbd,
	0x7c, 0xa7, 0x32, 0x97, 0x56, 0x3c, 0xbc, 0x88, 0xea, 0x82, 0xc5, 0xe0, 0x4f, 0x59, 0x6e, 0xd9,
	0xb5, 0x6d, 0x16, 0x03, 0xb5, 0x11, 0x1c, 0xa2, 0x56, 0xf6, 0xab, 0x13, 0x16, 0x81, 0xdf, 0xb0,
	0x69, 0x65, 0xbd, 0xdb, 0x45, 0x80, 0x56, 0x39, 0xb8, 0x8f, 0x9a, 0x0a, 0xb4, 0x4c, 0x55, 0x04,
	0xfe, 0xb4, 0x6d, 0xc7, 0xdd, 0x9b, 0x3f, 0x22, 0x75, 0x84, 0xf6, 0x5f, 0xee, 0xac, 0x66, 0xe1,
	0xa1, 0x25, 0x1d, 0xdf, 0x46, 0x33, 0x3a, 0xed, 0x14, 0x01, 0xbf, 0x69, 0x8b, 0x9b, 0x73, 0x82,
	0x99, 0xdd, 0x2a, 0x44, 0xcf, 0xe7, 0xe1, 0x3d, 0xd4, 0x4c, 0x35, 0xa8, 0x47, 0xe2, 0x40, 0xfa,
	0x2d, 0x5b, 0xe0, 0x7f, 0x17, 0x26, 0xfb, 0xc2, 0x67, 0x25, 0x2b, 0x6c, 0xdf, 0x65, 0x57, 0xc5,
	0x14, 0x1e, 0x5a, 0x92, 0x96, 0x3e, 0x78, 0xe8, 0x9f, 0x2b, 0x37, 0x00, 0xff, 0x8f, 0xa6, 0xd9,
	0x70, 0x28, 0x0f, 0x21, 0x1f, 0xea, 0x66, 0x7b, 0xd6, 0x61, 0xa6, 0xb7, 0x72, 0x37, 0x2d, 0xe2,
	0xf8, 0xc9, 0xc4, 0xca, 0xad, 0xfc, 0x5c, 0xe7, 0xdc, 0xaa, 0xa1, 0x6c, 0x3a, 0x29, 0xe8, 0x74,
	0x68, 0x8a, 0x35, 0x6b, 0x93, 0xa3, 0xb3, 0xa0, 0x76, 0x7c, 0x16, 0xd4, 0x4e, 0xce, 0x82, 0xda,
	0xbb, 0x71, 0xe0, 0x1d, 0x8d, 0x03, 0xef, 0x78, 0x1c, 0x78, 0x27, 0xe3, 0xc0, 0xfb, 0x3c, 0x0e,
	0xbc, 0xf7, 0x5f, 0x82, 0xda, 0xf3, 0x66, 0xb1, 0xc4, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x53, 0xac, 0x65, 0xf7, 0x05, 0x00, 0x00,
}
