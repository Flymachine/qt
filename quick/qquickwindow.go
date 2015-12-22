package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"unsafe"
)

type QQuickWindow struct {
	gui.QWindow
}

type QQuickWindow_ITF interface {
	gui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func PointerFromQQuickWindow(ptr QQuickWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func NewQQuickWindowFromPointer(ptr unsafe.Pointer) *QQuickWindow {
	var n = new(QQuickWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickWindow_") {
		n.SetObjectName("QQuickWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickWindow) QQuickWindow_PTR() *QQuickWindow {
	return ptr
}

//QQuickWindow::CreateTextureOption
type QQuickWindow__CreateTextureOption int64

const (
	QQuickWindow__TextureHasAlphaChannel = QQuickWindow__CreateTextureOption(0x0001)
	QQuickWindow__TextureHasMipmaps      = QQuickWindow__CreateTextureOption(0x0002)
	QQuickWindow__TextureOwnsGLTexture   = QQuickWindow__CreateTextureOption(0x0004)
	QQuickWindow__TextureCanUseAtlas     = QQuickWindow__CreateTextureOption(0x0008)
)

//QQuickWindow::RenderStage
type QQuickWindow__RenderStage int64

const (
	QQuickWindow__BeforeSynchronizingStage = QQuickWindow__RenderStage(0)
	QQuickWindow__AfterSynchronizingStage  = QQuickWindow__RenderStage(1)
	QQuickWindow__BeforeRenderingStage     = QQuickWindow__RenderStage(2)
	QQuickWindow__AfterRenderingStage      = QQuickWindow__RenderStage(3)
	QQuickWindow__AfterSwapStage           = QQuickWindow__RenderStage(4)
)

//QQuickWindow::SceneGraphError
type QQuickWindow__SceneGraphError int64

const (
	QQuickWindow__ContextNotAvailable = QQuickWindow__SceneGraphError(1)
)

func (ptr *QQuickWindow) ActiveFocusItem() *QQuickItem {
	defer qt.Recovering("QQuickWindow::activeFocusItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ActiveFocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) Color() *gui.QColor {
	defer qt.Recovering("QQuickWindow::color")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QQuickWindow_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {
	defer qt.Recovering("QQuickWindow::contentItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ContentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) SetColor(color gui.QColor_ITF) {
	defer qt.Recovering("QQuickWindow::setColor")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func NewQQuickWindow(parent gui.QWindow_ITF) *QQuickWindow {
	defer qt.Recovering("QQuickWindow::QQuickWindow")

	return NewQQuickWindowFromPointer(C.QQuickWindow_NewQQuickWindow(gui.PointerFromQWindow(parent)))
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {
	defer qt.Recovering("QQuickWindow::accessibleRoot")

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QQuickWindow_AccessibleRoot(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {
	defer qt.Recovering("connect QQuickWindow::activeFocusItemChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectActiveFocusItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeFocusItemChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {
	defer qt.Recovering("disconnect QQuickWindow::activeFocusItemChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectActiveFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeFocusItemChanged")
	}
}

//export callbackQQuickWindowActiveFocusItemChanged
func callbackQQuickWindowActiveFocusItemChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::activeFocusItemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeFocusItemChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {
	defer qt.Recovering("connect QQuickWindow::afterAnimating")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterAnimating(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterAnimating", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {
	defer qt.Recovering("disconnect QQuickWindow::afterAnimating")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterAnimating(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterAnimating")
	}
}

//export callbackQQuickWindowAfterAnimating
func callbackQQuickWindowAfterAnimating(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterAnimating")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterAnimating"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {
	defer qt.Recovering("connect QQuickWindow::afterRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {
	defer qt.Recovering("disconnect QQuickWindow::afterRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterRendering")
	}
}

//export callbackQQuickWindowAfterRendering
func callbackQQuickWindowAfterRendering(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterRendering")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {
	defer qt.Recovering("connect QQuickWindow::afterSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {
	defer qt.Recovering("disconnect QQuickWindow::afterSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterSynchronizing")
	}
}

//export callbackQQuickWindowAfterSynchronizing
func callbackQQuickWindowAfterSynchronizing(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterSynchronizing")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {
	defer qt.Recovering("connect QQuickWindow::beforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {
	defer qt.Recovering("disconnect QQuickWindow::beforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeRendering")
	}
}

//export callbackQQuickWindowBeforeRendering
func callbackQQuickWindowBeforeRendering(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::beforeRendering")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {
	defer qt.Recovering("connect QQuickWindow::beforeSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {
	defer qt.Recovering("disconnect QQuickWindow::beforeSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeSynchronizing")
	}
}

//export callbackQQuickWindowBeforeSynchronizing
func callbackQQuickWindowBeforeSynchronizing(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::beforeSynchronizing")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {
	defer qt.Recovering("QQuickWindow::clearBeforeRendering")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_ClearBeforeRendering(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) ConnectColorChanged(f func(v *gui.QColor)) {
	defer qt.Recovering("connect QQuickWindow::colorChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectColorChanged() {
	defer qt.Recovering("disconnect QQuickWindow::colorChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorChanged")
	}
}

//export callbackQQuickWindowColorChanged
func callbackQQuickWindowColorChanged(ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::colorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(v))
	}

}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImage_ITF) *QSGTexture {
	defer qt.Recovering("QQuickWindow::createTextureFromImage")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage2(ptr.Pointer(), gui.PointerFromQImage(image)))
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImage_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	defer qt.Recovering("QQuickWindow::createTextureFromImage")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.int(options)))
	}
	return nil
}

func (ptr *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	defer qt.Recovering("QQuickWindow::effectiveDevicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QQuickWindow_EffectiveDevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWindow) ConnectExposeEvent(f func(v *gui.QExposeEvent)) {
	defer qt.Recovering("connect QQuickWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "exposeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectExposeEvent() {
	defer qt.Recovering("disconnect QQuickWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "exposeEvent")
	}
}

//export callbackQQuickWindowExposeEvent
func callbackQQuickWindowExposeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickWindowFocusInEvent
func callbackQQuickWindowFocusInEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickWindowFocusOutEvent
func callbackQQuickWindowFocusOutEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {
	defer qt.Recovering("connect QQuickWindow::frameSwapped")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectFrameSwapped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameSwapped", f)
	}
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {
	defer qt.Recovering("disconnect QQuickWindow::frameSwapped")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectFrameSwapped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameSwapped")
	}
}

//export callbackQQuickWindowFrameSwapped
func callbackQQuickWindowFrameSwapped(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::frameSwapped")

	if signal := qt.GetSignal(C.GoString(ptrName), "frameSwapped"); signal != nil {
		signal.(func())()
	}

}

func QQuickWindow_HasDefaultAlphaBuffer() bool {
	defer qt.Recovering("QQuickWindow::hasDefaultAlphaBuffer")

	return C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer() != 0
}

func (ptr *QQuickWindow) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QQuickWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QQuickWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQQuickWindowHideEvent
func callbackQQuickWindowHideEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QQuickWindow) IncubationController() *qml.QQmlIncubationController {
	defer qt.Recovering("QQuickWindow::incubationController")

	if ptr.Pointer() != nil {
		return qml.NewQQmlIncubationControllerFromPointer(C.QQuickWindow_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {
	defer qt.Recovering("QQuickWindow::isPersistentOpenGLContext")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentOpenGLContext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {
	defer qt.Recovering("QQuickWindow::isPersistentSceneGraph")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentSceneGraph(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {
	defer qt.Recovering("QQuickWindow::isSceneGraphInitialized")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsSceneGraphInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickWindowKeyPressEvent
func callbackQQuickWindowKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickWindowKeyReleaseEvent
func callbackQQuickWindowKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickWindowMouseDoubleClickEvent
func callbackQQuickWindowMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	defer qt.Recovering("QQuickWindow::mouseGrabberItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_MouseGrabberItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickWindowMouseMoveEvent
func callbackQQuickWindowMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickWindowMousePressEvent
func callbackQQuickWindowMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickWindowMouseReleaseEvent
func callbackQQuickWindowMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	defer qt.Recovering("QQuickWindow::openglContext")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLContextFromPointer(C.QQuickWindow_OpenglContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context *gui.QOpenGLContext)) {
	defer qt.Recovering("connect QQuickWindow::openglContextCreated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectOpenglContextCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "openglContextCreated", f)
	}
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {
	defer qt.Recovering("disconnect QQuickWindow::openglContextCreated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectOpenglContextCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "openglContextCreated")
	}
}

//export callbackQQuickWindowOpenglContextCreated
func callbackQQuickWindowOpenglContextCreated(ptrName *C.char, context unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::openglContextCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "openglContextCreated"); signal != nil {
		signal.(func(*gui.QOpenGLContext))(gui.NewQOpenGLContextFromPointer(context))
	}

}

func (ptr *QQuickWindow) ReleaseResources() {
	defer qt.Recovering("QQuickWindow::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {
	defer qt.Recovering("QQuickWindow::renderTarget")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLFramebufferObjectFromPointer(C.QQuickWindow_RenderTarget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) RenderTargetSize() *core.QSize {
	defer qt.Recovering("QQuickWindow::renderTargetSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWindow_RenderTargetSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ResetOpenGLState() {
	defer qt.Recovering("QQuickWindow::resetOpenGLState")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ResetOpenGLState(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectResizeEvent(f func(ev *gui.QResizeEvent)) {
	defer qt.Recovering("connect QQuickWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QQuickWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQQuickWindowResizeEvent
func callbackQQuickWindowResizeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphAboutToStop")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphAboutToStop(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphAboutToStop")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphAboutToStop(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop")
	}
}

//export callbackQQuickWindowSceneGraphAboutToStop
func callbackQQuickWindowSceneGraphAboutToStop(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphAboutToStop")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphAboutToStop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWindowSceneGraphError
func callbackQQuickWindowSceneGraphError(ptrName *C.char, error C.int, message *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphInitialized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInitialized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInitialized", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphInitialized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInitialized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInitialized")
	}
}

//export callbackQQuickWindowSceneGraphInitialized
func callbackQQuickWindowSceneGraphInitialized(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphInitialized")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphInitialized"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphInvalidated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInvalidated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInvalidated", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphInvalidated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInvalidated")
	}
}

//export callbackQQuickWindowSceneGraphInvalidated
func callbackQQuickWindowSceneGraphInvalidated(ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphInvalidated")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ScheduleRenderJob(job core.QRunnable_ITF, stage QQuickWindow__RenderStage) {
	defer qt.Recovering("QQuickWindow::scheduleRenderJob")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ScheduleRenderJob(ptr.Pointer(), core.PointerFromQRunnable(job), C.int(stage))
	}
}

func (ptr *QQuickWindow) SendEvent(item QQuickItem_ITF, e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWindow::sendEvent")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_SendEvent(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	defer qt.Recovering("QQuickWindow::setClearBeforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetClearBeforeRendering(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	defer qt.Recovering("QQuickWindow::setDefaultAlphaBuffer")

	C.QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(C.int(qt.GoBoolToInt(useAlpha)))
}

func (ptr *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	defer qt.Recovering("QQuickWindow::setPersistentOpenGLContext")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentOpenGLContext(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	defer qt.Recovering("QQuickWindow::setPersistentSceneGraph")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentSceneGraph(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetRenderTarget(fbo gui.QOpenGLFramebufferObject_ITF) {
	defer qt.Recovering("QQuickWindow::setRenderTarget")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetRenderTarget(ptr.Pointer(), gui.PointerFromQOpenGLFramebufferObject(fbo))
	}
}

func (ptr *QQuickWindow) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QQuickWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QQuickWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQQuickWindowShowEvent
func callbackQQuickWindowShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QQuickWindow) Update() {
	defer qt.Recovering("QQuickWindow::update")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickWindowWheelEvent
func callbackQQuickWindowWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	defer qt.Recovering("QQuickWindow::~QQuickWindow")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
