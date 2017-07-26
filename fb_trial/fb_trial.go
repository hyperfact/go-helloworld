package main

import (
	"fmt"
	"runtime"

	sample "fb_trial/MyGame/Sample"

	fbs "github.com/google/flatbuffers/go"
)

func main() {
	fmt.Println(runtime.GOROOT())
	var fb fbs.FlatBuffer
	_ = fb

	builder := fbs.NewBuilder(1024)

	swordName := builder.CreateString("sword")
	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, swordName)
	sample.WeaponAddDamage(builder, 10)
	sword := sample.WeaponEnd(builder)

	builder.Finish(sword)
	data := builder.FinishedBytes()
	fmt.Println(data)

	wp := sample.GetRootAsWeapon(data, 0)
	fmt.Println(string(wp.Name()))
	fmt.Println(wp.Damage())
}
