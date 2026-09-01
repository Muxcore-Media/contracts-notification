package notificationv1

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestChannelEnumNumbersStable(t *testing.T) {
	want := map[string]int32{
		"CHANNEL_UNSPECIFIED": 0,
		"CHANNEL_DISCORD":     1,
		"CHANNEL_SLACK":       2,
		"CHANNEL_WEBHOOK":     3,
		"CHANNEL_EMAIL":       4,
		"CHANNEL_APPRISE":     5,
	}
	for name, num := range want {
		if v, ok := Channel_value[name]; !ok || v != num {
			t.Errorf("Channel %s: got %d want %d", name, v, num)
		}
	}
}

func TestNotificationServiceRPCNamesStable(t *testing.T) {
	want := []string{
		NotificationService_Notify_FullMethodName,
		NotificationService_Configure_FullMethodName,
		NotificationService_Status_FullMethodName,
		NotificationService_Test_FullMethodName,
	}
	expected := map[string]bool{
		"/muxcore.notification.v1.NotificationService/Notify":    true,
		"/muxcore.notification.v1.NotificationService/Configure": true,
		"/muxcore.notification.v1.NotificationService/Status":    true,
		"/muxcore.notification.v1.NotificationService/Test":      true,
	}
	for _, method := range want {
		if !expected[method] {
			t.Errorf("unexpected full method name constant %q", method)
		}
	}
	got := make(map[string]bool, len(NotificationService_ServiceDesc.Methods))
	for _, m := range NotificationService_ServiceDesc.Methods {
		got["/muxcore.notification.v1.NotificationService/"+m.MethodName] = true
	}
	for name := range expected {
		if !got[name] {
			t.Errorf("missing RPC %s", name)
		}
	}
}

func TestNotifyRequestFieldTagsStable(t *testing.T) {
	msg := (&NotifyRequest{}).ProtoReflect().Descriptor()
	want := map[string]protoreflect.FieldNumber{
		"title":         1,
		"message":       2,
		"severity":      3,
		"source_module": 4,
		"channels":      5,
		"fields":        6,
	}
	for name, num := range want {
		f := msg.Fields().ByName(protoreflect.Name(name))
		if f == nil {
			t.Fatalf("field %q not found", name)
		}
		if f.Number() != num {
			t.Errorf("NotifyRequest.%s: got field %d want %d", name, f.Number(), num)
		}
	}
}

func TestSeverityEnumNumbersStable(t *testing.T) {
	want := map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_INFO":        1,
		"SEVERITY_SUCCESS":     2,
		"SEVERITY_WARNING":     3,
		"SEVERITY_ERROR":       4,
	}
	for name, num := range want {
		if v, ok := Severity_value[name]; !ok || v != num {
			t.Errorf("Severity %s: got %d want %d", name, v, num)
		}
	}
}

func TestImportPath(t *testing.T) {
	var _ grpc.ServiceDesc = NotificationService_ServiceDesc
	if NotificationService_ServiceDesc.ServiceName != "muxcore.notification.v1.NotificationService" {
		t.Fatalf("service name drift: %s", NotificationService_ServiceDesc.ServiceName)
	}
}

func TestCapabilityIDs(t *testing.T) {
	const consumerCapability = "notification"
	const contractCapability = "contracts.notification"
	if consumerCapability == contractCapability {
		t.Fatal("consumer and contract capability ids must differ")
	}
}

func TestNotifyRequestSeverityType(t *testing.T) {
	f := (&NotifyRequest{}).ProtoReflect().Descriptor().Fields().ByName("severity")
	if f == nil {
		t.Fatal("severity field not found")
	}
	if f.Kind() != protoreflect.EnumKind {
		t.Fatalf("severity kind: got %v want EnumKind", f.Kind())
	}
	if f.Enum().FullName() != "muxcore.notification.v1.Severity" {
		t.Fatalf("severity enum: got %s", f.Enum().FullName())
	}
}
