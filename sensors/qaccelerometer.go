package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccelerometer struct {
	QSensor
}

type QAccelerometer_ITF interface {
	QSensor_ITF
	QAccelerometer_PTR() *QAccelerometer
}

func PointerFromQAccelerometer(ptr QAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometer_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerFromPointer(ptr unsafe.Pointer) *QAccelerometer {
	var n = new(QAccelerometer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAccelerometer_") {
		n.SetObjectName("QAccelerometer_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccelerometer) QAccelerometer_PTR() *QAccelerometer {
	return ptr
}

//QAccelerometer::AccelerationMode
type QAccelerometer__AccelerationMode int64

const (
	QAccelerometer__Combined = QAccelerometer__AccelerationMode(0)
	QAccelerometer__Gravity  = QAccelerometer__AccelerationMode(1)
	QAccelerometer__User     = QAccelerometer__AccelerationMode(2)
)

func (ptr *QAccelerometer) AccelerationMode() QAccelerometer__AccelerationMode {
	defer qt.Recovering("QAccelerometer::accelerationMode")

	if ptr.Pointer() != nil {
		return QAccelerometer__AccelerationMode(C.QAccelerometer_AccelerationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometer) Reading() *QAccelerometerReading {
	defer qt.Recovering("QAccelerometer::reading")

	if ptr.Pointer() != nil {
		return NewQAccelerometerReadingFromPointer(C.QAccelerometer_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQAccelerometer(parent core.QObject_ITF) *QAccelerometer {
	defer qt.Recovering("QAccelerometer::QAccelerometer")

	return NewQAccelerometerFromPointer(C.QAccelerometer_NewQAccelerometer(core.PointerFromQObject(parent)))
}

func (ptr *QAccelerometer) ConnectAccelerationModeChanged(f func(accelerationMode QAccelerometer__AccelerationMode)) {
	defer qt.Recovering("connect QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_ConnectAccelerationModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accelerationModeChanged", f)
	}
}

func (ptr *QAccelerometer) DisconnectAccelerationModeChanged() {
	defer qt.Recovering("disconnect QAccelerometer::accelerationModeChanged")

	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectAccelerationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accelerationModeChanged")
	}
}

//export callbackQAccelerometerAccelerationModeChanged
func callbackQAccelerometerAccelerationModeChanged(ptrName *C.char, accelerationMode C.int) {
	defer qt.Recovering("callback QAccelerometer::accelerationModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "accelerationModeChanged"); signal != nil {
		signal.(func(QAccelerometer__AccelerationMode))(QAccelerometer__AccelerationMode(accelerationMode))
	}

}

func (ptr *QAccelerometer) SetAccelerationMode(accelerationMode QAccelerometer__AccelerationMode) {
	defer qt.Recovering("QAccelerometer::setAccelerationMode")

	if ptr.Pointer() != nil {
		C.QAccelerometer_SetAccelerationMode(ptr.Pointer(), C.int(accelerationMode))
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometer() {
	defer qt.Recovering("QAccelerometer::~QAccelerometer")

	if ptr.Pointer() != nil {
		C.QAccelerometer_DestroyQAccelerometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
