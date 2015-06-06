package chess

const (
	_ = iota
	ROBOT
	PEOPLE
)

type Army struct {
	Panel *Panel
	Role  int
	win   bool
}

func NewArmy(panel *Panel, role int) *Army {
	return &Army{
		Panel: panel, Role: role, win: false,
	}
}

func (army *Army) IsWin() bool {
	return army.win
}

func (army *Army) DownPiece(i, j int) bool {
	if army.Panel.Panel[i][j] == 0 {
		return false
	}
	army.Panel.Panel[i][j] = army.Role
	army.win = NewPiece(army.Panel, i, j).IsWin()
	return true
}

type Panel struct {
	Rows, Columns int
	Panel         [][]int
}

func NewPanel(rows, columns int) *Panel {
	var panel [][]int = make([][]int, rows, rows)
	for i := 0; i < rows; i++ {
		panel[i] = make([]int, columns, columns)
	}
	return &Panel{
		Rows: rows, Columns: columns, Panel: panel,
	}
}

func (panel *Panel) IsFull() bool {
	for _, array := range panel.Panel {
		for _, v := range array {
			if v == 0 {
				return false
			}
		}
	}
	return true
}

func (panel *Panel) Draw() {

}

type Piece struct {
	Panel *Panel
	Army  int
	X, Y  int
}

func NewPiece(panel *Panel, x int, y int) *Piece {
	return &Piece{
		Panel: panel, X: x, Y: y,
	}
}

func (piece *Piece) IsWin() bool {
	//水平
	if piece.Quantity(0, 1)+piece.Quantity(0, -1) >= 4 {
		return true
	}
	//垂直
	if piece.Quantity(1, 0)+piece.Quantity(-1, 0) >= 4 {
		return true
	}
	//左上右下斜
	if piece.Quantity(-1, -1)+piece.Quantity(1, 1) >= 4 {
		return true
	}
	//左下右上斜
	if piece.Quantity(-1, 1)+piece.Quantity(1, -1) >= 4 {
		return true
	}
	return false
}

func (piece *Piece) Quantity(x, y int) int {
	var counter int
	for row, column := piece.X, piece.Y; row >= 0 && column >= 0 && row < piece.Panel.Rows && column < piece.Panel.Columns; {
		if 0 <= row+x && row+x < piece.Panel.Rows && 0 <= column+y && column+y < piece.Panel.Columns && piece.Panel.Panel[row+x][column+y] == piece.Army {
			counter++
		} else {
			break
		}
		row++
		column++
	}
	return counter
}

type Game struct {
	Panel         *Panel
	Robot, People *Army
	currentArmy   *Army
}

func NewGame(rows, columns int) *Game {
	panel := NewPanel(rows, columns)
	robot, people := NewArmy(panel, ROBOT), NewArmy(panel, PEOPLE)
	return &Game{
		Panel:       panel,
		Robot:       robot,
		People:      people,
		currentArmy: people,
	}
}

func (game *Game) changeArmy() *Army {
	if game.currentArmy == game.People {
		game.currentArmy = game.Robot
	} else {
		game.currentArmy = game.People
	}
	return game.currentArmy
}

func (game *Game) Run() {
}
