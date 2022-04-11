package main

//Исполняемоый пакет для работы с консольными флагами
//В целом, на него смотреть пока не нужно, это я сам добью, потом посмотрите и оцените(!)
import (
	"DMTI_baza/natural"
	"DMTI_baza/polynome"
	"DMTI_baza/rational"
	"DMTI_baza/whole"
	"fmt"
	"strconv"
	"strings"
)

//Модуль от Голубева Михаила
func main() {
	var moduleNumber int
	var functionNumber int
	fmt.Println("введите номер модуля")
	fmt.Scan(&moduleNumber)
	fmt.Println("введите номер функции")
	fmt.Scan(&functionNumber)
	switch moduleNumber {
	case 1:
		naturalsGui(functionNumber)
	case 2:
		wholeGui(functionNumber)
	case 3:
		rationalGui(functionNumber)
	case 4:
		polynomeGui(functionNumber)
	}
}
func naturalsGui(f int) {
	switch f {
	case 1: //сравнение
		fmt.Println("Это функция Сравнение натуральных чисел: 2 - если первое больше второго, 0, если равно, 1 иначе.")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.Compare(n1, n2))
	case 2: //проверка на 0
		fmt.Println("Это функция Проверка на ноль: если число не равно нулю, то «да» иначе «нет»")
		fmt.Println("Натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println(natural.CheckNull(n1))
	case 3: //доавление 1
		fmt.Println("Это функция  Добавление 1 к натуральному числу")
		fmt.Println("Натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println(natural.ToStringN(natural.Addition1(n1)))
	case 4: //сложение
		fmt.Println("Это функция Сложение натуральных чисел ")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.ToStringN(natural.Addition(n1, n2)))
	case 5: //Вычитание из первого большего натурального числа второго меньшего или равного
		fmt.Println("Это функция Вычитание из первого большего натурального числа второго меньшего или равного")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.ToStringN(natural.Subtraction(n1, n2)))
	case 6: //умножение на цифру
		fmt.Println("Это функция Умножение натурального числа на цифру")
		fmt.Println("Первое натуральное число - ")
		var s string
		var k uint8
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&k)
		fmt.Println(natural.ToStringN(natural.MultiplicationNaturalNumber(n1, k)))
	case 7: //Умножение натурального числа на 10^k
		fmt.Println("Это функция Умножение натурального числа на 10^k")
		fmt.Println("Первое натуральное число - ")
		var s string
		var k int
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&k)
		fmt.Println(natural.ToStringN(natural.MultiplicationBy10k(n1, k)))
	case 8: //Умножение натуральных чисел
		fmt.Println("Это функция Умножение натуральных чисел")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.Multiplication(n1, n2))
	case 9: //Вычитание из натурального другого натурального, умноженного на цифру для случая с неотрицательным результатом
		fmt.Println("Это функция Вычитание из натурального другого натурального, умноженного на цифру для случая с неотрицательным результатом")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		fmt.Println("k=")
		var k uint8
		fmt.Scan(&k)
		n2 := inputNatural(s)
		fmt.Println(natural.DifferenceOfNaturals(n1, n2, k))
	case 10: //Вычисление первой цифры деления большего натурального на меньшее, домноженное на 10^k,где k - номер позиции этой цифры (номер считается с нуля)
		fmt.Println("Это функция Вычисление первой цифры деления большего натурального на меньшее, домноженное на 10^k,где k - номер позиции этой цифры (номер считается с нуля)")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.DivideOneIteration(n1, n2))
	case 11: //Частное от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля)
		fmt.Println("Это функция Частное от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля)")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.IntegerFromDivision(n1, n2))
	case 12: //Остаток от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля)
		fmt.Println("Это функция Остаток от деления большего натурального числа на меньшее или равное натуральное с остатком(делитель отличен от нуля)")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.RemainderFromDivision(n1, n2))
	case 13: //нод
		fmt.Println("Это функция НОД натуральных чисел")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.GreatestCommonDivisor(n1, n2))
	case 14: //НОК
		fmt.Println("Это функция НОК натуральных чисел")
		fmt.Println("Первое натуральное число - ")
		var s string
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println("Второе натуральное число - ")
		fmt.Scan(&s)
		n2 := inputNatural(s)
		fmt.Println(natural.LeastCommonMultiple(n1, n2))
	}
}

func inputNatural(s string) natural.Natural {
	nat := strings.Split(s, "")
	var natur []uint8
	for _, v := range nat {
		n, _ := strconv.Atoi(v)
		natur = append(natur, uint8(n))
	}
	var n natural.Natural
	n.MakeN(natur)
	return n
}

func wholeGui(f int) {
	var s string
	switch f {
	case 1: //Абсолютная величина числа, результат - натуральное
		fmt.Println("Это функция Абсолютная величина числа, результат - натуральное")
		fmt.Println("Целое число-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(natural.ToStringN(whole.Absolute(w1)))
	case 2: //Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное)
		fmt.Println("Это функция Определение положительности числа (2 - положительное, 0 — равное нулю, 1 - отрицательное)")
		fmt.Println("Целое число-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(whole.Positivity(w1))
	case 3: //Умножение целого на (-1)
		fmt.Println("Это функция Умножение целого на (-1)")
		fmt.Println("Целое число-")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.MultiplicationByNegativeOne(w1)))
	case 4: //Преобразование натурального в целое
		fmt.Println("Это функция Преобразование натурального в целое")
		fmt.Println("Натуральное -")
		fmt.Scan(&s)
		n1 := inputNatural(s)
		fmt.Println(whole.ToStringW(whole.FromNaturalsToWhole(n1)))
	case 5: //Преобразование целого неотрицательного в натуральное
		fmt.Println("Это функция Преобразование целого неотрицательного в натуральное")
		fmt.Println("Целое неотрицательное  -")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println(natural.ToStringN(whole.FromWholeToNaturals(w1)))
	case 6: //Сложение целых чисел
		fmt.Println("Это функция Сложение целых чисел")
		fmt.Println("Первое слагаемое - ")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("Второе слагаемое - ")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.Addition(w1, w2)))
	case 7: //вычитание целых чисел
		fmt.Println("Это функция Вычитание целых чисел")
		fmt.Println("Откуда вычитаем - ")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("Что вычитаем - ")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.Subtraction(w1, w2)))
	case 8: //умножение целых чисел
		fmt.Println("Это функция умножение целых чисел")
		fmt.Println("Первый множитель - ")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("Второй множитель - ")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.Multiplication(w1, w2)))
	case 9: //Частное от деления целого на целое (делитель отличен от нуля)
		fmt.Println("Это функция Частное от деления целого на целое (делитель отличен от нуля)")
		fmt.Println("Первое целое - ")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("Второе целое - ")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.WholeFromDivision(w1, w2)))
	case 10: //Остаток от деления целого на целое(делитель отличен от нуля)
		fmt.Println("Это функция Остаток от деления целого на целое(делитель отличен от нуля)")
		fmt.Println("Делимое - ")
		fmt.Scan(&s)
		w1 := inputWhole(s)
		fmt.Println("Делитель - ")
		fmt.Scan(&s)
		w2 := inputWhole(s)
		fmt.Println(whole.ToStringW(whole.RemainderFromDivision(w1, w2)))
	}
}

func inputWhole(s string) whole.Whole {
	var negative = false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	nums := inputNatural(s).Digits
	var w whole.Whole
	w.MakeW(negative, nums)
	return w
}

func rationalGui(f int) {
	var s string
	switch f {
	case 1: //Сокращение дроби
		fmt.Println("Это функция Сокращение дроби")
		fmt.Println("Дробь - ")
		fmt.Scan(&s)
		r := inputRational(s)
		fmt.Println(rational.ToStringR(rational.SimplifyingFractions(r)))
	case 2: //Проверка на целое, если рациональное число является целым, то «да», иначе «нет»
		fmt.Println("Это функция Проверка на целое, если рациональное число является целым, то «да», иначе «нет»")
		fmt.Println("Дробь - ")
		fmt.Scan(&s)
		r := inputRational(s)
		fmt.Println(rational.CheckingForWhole(r))
	case 3: //Преобразование целого в дробное
		fmt.Println("Это функция Преобразование целого в дробное")
		fmt.Println("Целое - ")
		fmt.Scan(&s)
		wh := inputWhole(s)
		fmt.Println(rational.ToStringR(rational.WholeToFractional(wh)))
	case 4: //дробного в целое
		fmt.Println("Это функция Преобразование дробного в целое (если знаменатель равен 1)")
		fmt.Println("Дробное - ")
		fmt.Scan(&s)
		r := inputRational(s)
		fmt.Println(whole.ToStringW(rational.FractionalToWhole(r)))
	case 5: //сложение дробей
		fmt.Println("Это функция Сложение дробей")
		fmt.Println("Первое слагаемое - ")
		fmt.Scan(&s)
		r1 := inputRational(s)
		fmt.Println("Второе слагаемое - ")
		fmt.Scan(&s)
		r2 := inputRational(s)
		fmt.Println(rational.ToStringR(rational.Addition(r1, r2)))
	case 6: //вычитание дробей
		fmt.Println("Это функция вычитание дробей")
		fmt.Println("Уменьшаемое - ")
		fmt.Scan(&s)
		r1 := inputRational(s)
		fmt.Println("Вычитаемое - ")
		fmt.Scan(&s)
		r2 := inputRational(s)
		fmt.Println(rational.ToStringR(rational.Subtraction(r1, r2)))
	case 7: //умножение дробей
		fmt.Println("Это функция умножение дробей")
		fmt.Println("Первый множитель - ")
		fmt.Scan(&s)
		r1 := inputRational(s)
		fmt.Println("Второй множитель - ")
		fmt.Scan(&s)
		r2 := inputRational(s)
		fmt.Println(rational.ToStringR(rational.Multiplication(r1, r2)))
	case 8: //деление дробей
		fmt.Println("Это функция деление дробей")
		fmt.Println("Делимое - ")
		fmt.Scan(&s)
		r1 := inputRational(s)
		fmt.Println("Делитель - ")
		fmt.Scan(&s)
		r2 := inputRational(s)
		fmt.Println(rational.ToStringR(rational.Division(r1, r2)))
	}
}

func inputRational(s string) rational.Rational {
	arr := strings.Split(s, "/")
	nom := inputWhole(arr[0])
	denom := inputNatural(arr[1])
	var r rational.Rational
	r.MakeR(nom, denom)
	return r
}

func polynomeGui(f int) {
	var s, sTemp string
	var power int
	switch f {
	case 1: //Сложение многочленов
		fmt.Println("Это функция Сложение многочленов")
		fmt.Println("Степень первого многочлена")
		fmt.Scan(&power)
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		p1 := inputPolynomes(s)
		s = ""
		fmt.Println("Степень сторого многочлена - ")
		var pw int
		fmt.Scan(&pw)
		for i := 0; i <= pw; i++ {
			fmt.Println("Коэффициент при степени ", pw-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < pw {
				s += ";"
			}
		}
		p2 := inputPolynomes(s)
		res := polynome.AdditionP(p1, p2)
		fmt.Println(res.ToStringPol())
	case 2: //вычитание многочленов
		fmt.Println("Это функция вычитание многочленов")
		fmt.Println("Степень уменьшаемого многочлена - ")
		fmt.Scan(&power)
		s = ""
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		from := inputPolynomes(s)
		fmt.Println("Степень вычитаемого многочлена - ")
		fmt.Scan(&power)
		s = ""
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		what := inputPolynomes(s)
		res := polynome.SubstractionP(from, what)
		fmt.Println(res.ToStringPol())
	case 3: //Умножение многочлена на рациональное число
		fmt.Println("Это функция Умножение многочлена на рациональное число")
		fmt.Println("Степень многочлена = ")
		fmt.Scan(&power)
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		p1 := inputPolynomes(s)
		fmt.Println("Рациональное число множитель - ")
		s = ""
		fmt.Scan(&s)
		r := inputRational(s)
		res := polynome.MultiplicationRational(p1, r)
		fmt.Println(res.ToStringPol())
	case 4: //Умножение многочлена на x^k
		fmt.Println("Это функция Умножение многочлена на x^k")
		fmt.Println("Степень многочлена = ")
		fmt.Scan(&power)
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		p1 := inputPolynomes(s)
		fmt.Println("Степень k=")
		s = ""
		fmt.Scan(&power)
		res := polynome.MultiplicationXpowerK(p1, power)
		fmt.Println(res.ToStringPol())
	case 5: //старший коэфф
		fmt.Println("Это функция Старший коэффициент многочлена")
		fmt.Println("Степень многочлена = ")
		fmt.Scan(&power)
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		p1 := inputPolynomes(s)
		fmt.Println(rational.ToStringR(polynome.OlderCoeffPoly(p1)))
	case 6: //степень
		fmt.Println("Это функция Степень многочлена")
		fmt.Println("Степень многочлена = ")
		fmt.Scan(&power)
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		p1 := inputPolynomes(s)
		fmt.Println(polynome.OlderPoly(p1))
	case 12: //производная
		fmt.Println("Это функция Производная многочлена")
		fmt.Println("Степень многочлена = ")
		fmt.Scan(&power)
		for i := 0; i <= power; i++ {
			fmt.Println("Коэффициент при степени ", power-i, ":")
			fmt.Scan(&sTemp)
			s += sTemp
			if i < power {
				s += ";"
			}
		}
		p1 := inputPolynomes(s)
		res := polynome.Derivative(p1)
		fmt.Println(res.ToStringPol())
	}
}

func inputPolynomes(s string) polynome.Polynomial {
	arrayRationalsStr := strings.Split(s, ";")
	var coeffs []rational.Rational
	for i, _ := range arrayRationalsStr {
		coeffs = append(coeffs, inputRational(arrayRationalsStr[i]))
	}
	var p polynome.Polynomial
	p.MakeP(coeffs)
	return p
}
