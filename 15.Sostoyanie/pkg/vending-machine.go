package pkg

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	CurrentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
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
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}
	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.CurrentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.CurrentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.CurrentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.CurrentState.DispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.CurrentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	v.itemCount += count
}
