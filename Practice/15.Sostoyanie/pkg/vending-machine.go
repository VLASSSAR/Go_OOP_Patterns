package pkg

type VendingMachine struct {
	hasMoney      State
	hasItem       State
	noItem        State
	itemRequested State
	currentState  State
	itemPrice     int
	itemCount     int
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	v.itemCount += count
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func NewVendingMachine(itemCount int, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.noItem = noItemState
	v.hasMoney = hasMoneyState
	v.itemRequested = itemRequestedState
	return v
}
