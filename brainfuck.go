package main

import "fmt"

type Interpreter struct {
	cell       []byte
	pointer    int
	code       string
	codelen    int
	codepos    int
	maxpointer int
}

func (self *Interpreter) CellValues() {
	str := ""
	for i := 0; i <= self.maxpointer; i++ {
		f := "%d "
		if i == self.pointer {
			f = "[%d] "
		}
		str += fmt.Sprintf(f, self.cell[i])
	}
	fmt.Printf("\n%s\n", str)
}

func (self *Interpreter) Start(code string) {
	self.code = code
	self.pointer = 0
	self.codelen = len(code)
	self.cell = make([]byte, 500)
}

func (self *Interpreter) Run(cv bool) {
	if cv {
		defer self.CellValues()
	}

	for self.codepos < self.codelen {
		switch self.code[self.codepos] {
		case '>':
			{
				if self.pointer < 499 {
					self.pointer++
					if self.pointer > self.maxpointer {
						self.maxpointer = self.pointer
					}
				} else {
					fmt.Printf("Estouro: %d\n", self.codepos)
					return
				}
				break
			}
		case '<':
			{
				if self.pointer == 0 {
					fmt.Printf("Celula negativa: %d\n", self.codepos)
					return
				}
				self.pointer--
				break
			}
		case '+':
			self.cell[self.pointer]++
			break
		case '-':
			self.cell[self.pointer]--
			break
		case '.':
			fmt.Printf("%c", self.cell[self.pointer])
			break
		case '[':
			{
				f := false
				for i := self.codepos + 1; i < self.codelen; i++ {
					if self.code[i] == ']' {
						f = true
						break
					}
				}
				if !f {
					fmt.Printf("Não foi possível encontrar o operador \"%s\"\n", "]")
					return
				}
			}
		case ']':
			{
				if self.cell[self.pointer] != 0 {
					for i := self.codepos; i > 0; i-- {
						if self.code[i] == '[' {
							self.codepos = i - 1
							break
						}
					}
				}
			}
		case ',':
			{
				_, err := fmt.Scanf("%c", &self.cell[self.pointer])
				if err != nil {
					fmt.Errorf("%s\n", err)
				}
			}
		}
		self.codepos++
	}
}

func main() {
	f := Interpreter{}
	f.Start(`++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.`)
	f.Run(true)
}