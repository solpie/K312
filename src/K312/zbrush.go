package main

func setup() {
	keyStates = make(map[int]bool)
	var kfc *KFC
	var zb App
	zb.title = "ZBrush"

	kfc = zb.addDownKFC(VK_3)
	kfc.addKFC(rmbDown, 0)
	kfc = zb.addUpKFC(VK_3)
	kfc.addKFC(rmbUp, 5)

	kfc = zb.addDownKFC(VK_2)
	kfc.addKFC(keyDown, VK_CTRL)
	kfc.addKFC(rmbDown, 0)
	kfc = zb.addUpKFC(VK_2)
	kfc.addKFC(rmbUp, 5)
	kfc.addKFC(keyUp, VK_CTRL)

	kfc = zb.addDownKFC(VK_1)
	kfc.addKFC(keyDown, VK_ALT)
	kfc.addKFC(rmbDown, 0)
	kfc = zb.addUpKFC(VK_1)
	kfc.addKFC(rmbUp, 5)
	kfc.addKFC(keyUp, VK_ALT)

	apps = append(apps, zb)
}
