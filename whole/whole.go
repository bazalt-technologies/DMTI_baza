package whole

import "DMTI_baza/natural"

// Whole
// Структура целого числа
type Whole struct {
	Num      natural.Natural
	Negative bool // true (1), если отрицательное; false (0), если положительное
}

// Zero
// Возвращает целый нуль (положительный =) )
func Zero() Whole {
	var zero Whole
	zero.Num = natural.Zero()
	zero.Negative = false
	return zero
}

// ToStringW
// Возвращает целое число как строку
func ToStringW(w Whole) string {
	var s string
	if w.Negative {
		s += "-"
	}
	s += natural.ToStringN(w.Num)
	return s
}

// MakeW
// Метод для создания целого
func (w *Whole) MakeW(Negative bool, digits []uint8) {
	for _, v := range digits {
		if v != 0 { //Проверка на 0 старшей цифры
			break
		} else {
			digits = digits[1:]
		}
	}
	if len(digits) == 0 {
		digits = append(digits, 0)
	}
	w.Num.Digits = digits
	w.Num.Older = uint32(len(digits)) - 1
	w.Negative = Negative
}

// Absolute Тростина Максима
// Возвращает модуль целого числа как натуральное
func Absolute(w Whole) natural.Natural {
	var n natural.Natural
	n.MakeN(w.Num.Digits)
	return n
}

// Positivity Турбина
// Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное)
func Positivity(x Whole) int {
	switch {
	case x.Num.Digits[0] == 0:
		return 0
	case x.Negative == true:
		return 1
	default:
		return 2
	}
}

// MultiplicationByNegativeOne Хвостовский
// Умножение целого на (-1)
func MultiplicationByNegativeOne(x Whole) Whole {
	x.Negative = !x.Negative
	return x
}

// FromNaturalsToWhole Комаровский
// Преобразование из натурального в целое
func FromNaturalsToWhole(nat natural.Natural) Whole {
	var res Whole
	res.Num = nat
	res.Negative = false
	return res
}

// FromWholeToNaturals Комаровский
// Преобразование из неотрицательного целого в натуральное
func FromWholeToNaturals(wh Whole) natural.Natural {
	var res natural.Natural
	res = wh.Num
	return res

}

// Multiplication Тростин Максим
// Умножение целых
func Multiplication(x, y Whole) Whole {
	// Если хотя бы один ноль, то возвращаем ноль
	// (чтобы не проходить по всем функциям и не париться по поводу знака)
	if natural.CheckNull(x.Num) || natural.CheckNull(y.Num) {
		return Zero()
	}
	var res Whole
	// Определяем знак результата
	if (x.Negative && y.Negative) || (!x.Negative && !y.Negative) {
		res.Negative = false
	} else {
		res.Negative = true
	}
	// Вычисляем модуль результата
	res.Num = natural.Multiplication(x.Num, y.Num)
	return res
}
