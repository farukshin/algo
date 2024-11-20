# Система непересекающихся множеств

Disjoint set union (DSU) - структура данных которая позволяет проводить операции с непересекающимися множествами, объединять их, отвечать на определенные вопросы, например "находится ли элемент a в одном множестве с элементом b" или "чему размер размер множества с элементом a".

Обе операции выполняются за ассимптотику O(1) с небольшим коэффициентом.

## Принцип работы структуры

Создадим структуру и будем хранить в ней количество элементов n, массив лидеров для каждой вершины и массив количества элементов в множестве.

```
type DSU struct {
	n   int
	p   []int
	cnt []int
}
```

Вначале пусть все элементы множества (n) находятся в отдельных множествах. Поэтому при инициализации массив p будет заполняться так:

```
func (dsu *DSU) init() {
	for i := 0; i < dsu.n; i++ {
		dsu.p[i] = i
	}
}
```

Для запросов "в каком множестве находится элемент a" необходимо подняться по ссылкам до корня (лидера множества):

```
func (dsu *DSU) leader(v int) int {
	if dsu.p[v] == v {
		return v
	}	
	return dsu.leader(dsu.p[v])
}
```

Для объединения двух множеств нужно подвесить корень одного множества к другому:

```
func (dsu *DSU) unite(a int, b int) {
	dsu.p[a] = dsu.leader(b)
}
```

## Оптимизации

**Эвристика сжатия пути.** Перед тем как вернуть ответ, запишем его в элемент массива p текущего элемента

```
func (dsu *DSU) leader(v int) int {
	if dsu.p[v] == v {
		return v
	}
	dsu.p[v] = dsu.leader(dsu.p[v])
	return dsu.p[v]
}
```

**Весовая эвристика.** Будем для каждой вершины хранить размер текущего множества, при объединении подвешивать меньшее множество к большему:

```
func (dsu *DSU) unite(a int, b int) {
	a = dsu.leader(a)
	b = dsu.leader(b)
	if dsu.cnt[a] < dsu.cnt[b] {
		a, b = b, a
	}
	dsu.cnt[a] += dsu.cnt[b]
	dsu.p[b] = a
}
```

## Итоговая реализация

```
type DSU struct {
	n   int
	p   []int
	cnt []int
}

func (dsu *DSU) init() {
	for i := 0; i < dsu.n; i++ {
		dsu.p[i] = i
	}
}

func (dsu *DSU) leader(v int) int {
	if dsu.p[v] == v {
		return v
	}
	dsu.p[v] = dsu.leader(dsu.p[v])
	return dsu.p[v]
}

func (dsu *DSU) unite(a int, b int) {
	a = dsu.leader(a)
	b = dsu.leader(b)
	if dsu.cnt[a] < dsu.cnt[b] {
		a, b = b, a
	}
	dsu.cnt[a] += dsu.cnt[b]
	dsu.p[b] = a
}

func (dsu *DSU) uniteOld(a int, b int) {
	dsu.p[a] = dsu.leader(b)
}
```