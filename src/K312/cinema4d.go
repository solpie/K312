package main

func setupCinema4d() {
	var kfc *KFC
	var c4d App
	c4d.title = "CINEMA"

	kfc = c4d.addDownKFC(VK_3)
	kfc.addKFC(keyDown, VK_ALT)
	kfc.addKFC(lmbDown, 0)
	kfc = c4d.addUpKFC(VK_3)
	kfc.addKFC(lmbUp, 5)
	kfc.addKFC(keyUp, VK_ALT)

	kfc = c4d.addDownKFC(VK_2)
	kfc.addKFC(keyDown, VK_ALT)
	kfc.addKFC(rmbDown, 0)
	kfc = c4d.addUpKFC(VK_2)
	kfc.addKFC(rmbUp, 5)
	kfc.addKFC(keyUp, VK_ALT)

	kfc = c4d.addDownKFC(VK_1)
	kfc.addKFC(keyDown, VK_ALT)
	kfc.addKFC(mmbDown, 0)
	kfc = c4d.addUpKFC(VK_1)
	kfc.addKFC(mmbUp, 5)
	kfc.addKFC(keyUp, VK_ALT)

	apps = append(apps, c4d)
}
