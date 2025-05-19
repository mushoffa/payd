package valueobject

type ShiftRole int

const (
	RoleBarista ShiftRole = iota
	RoleCashier
	RoleWaitress
)

var shiftRoles = map[ShiftRole]string{
	RoleBarista:  "BARISTA",
	RoleCashier:  "CASHIER",
	RoleWaitress: "WAITRESS",
}

func (e ShiftRole) String() string {
	return shiftRoles[e]
}
