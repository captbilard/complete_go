package exercise

import (
	"fmt"
	"slices"
)

type Player struct {
	name      string
	inventory []Item
}

type Item struct {
	name string
	Type string
}

func (p *Player) PickUpItem(item Item) {
	p.inventory = append(p.inventory, item)
	fmt.Printf("Item %s of type %s has been added to your inventory\n", item.name, item.Type)
}

func (p *Player) DropItem(name string) {
	itemIndex := -1
	for i, item := range p.inventory {
		if item.name == name {
			itemIndex = i
			p.inventory = slices.Delete(p.inventory, itemIndex, itemIndex+1)
			fmt.Printf("%s has been removed from your inventory", name)
			return
		}
	}
	fmt.Printf("%s doesn't have %s in inventory", p.name, name)
}

func (p *Player) UseItem(itemName string) {
	for _, item := range p.inventory {
		if item.name == itemName {
			switch item.Type {
			case "potion":
				fmt.Println("Potion applied")
				p.DropItem(itemName)
			case "spell":
				fmt.Println("Avada Kedavra")
			case "map":
				fmt.Println("switched to Navigation mode")
			default:
				fmt.Printf("%s has been applied successfully", itemName)
			}
			return
		}
	}
	fmt.Printf("%s doesn't have %s in inventory", p.name, itemName)
}
