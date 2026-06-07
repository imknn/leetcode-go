package other

import "sort"

// LCP 82. 万灵之树
// 算法：枚举所有满二叉树形状 → 拆成"块"(B0..Bn) → DFS 枚举每侧排列 → meet-in-the-middle 计数
// 参考题解：https://leetcode.cn/problems/cnHoX6/solutions/3748723/mo-ling-zhi-shu-ti-jie-by-gpt-5-thinking-aplb/
// 原作者：西西弗斯
// 时间 O(Catalan(n) · 2^n · n)，空间 O(2^n)

// --- 轻量结构 ---

// intList 动态 int 数组
type intList struct {
	a   []int
	sz  int
}

func newIntList(cap int) *intList {
	return &intList{a: make([]int, cap)}
}

func (l *intList) add(v int) {
	if l.sz == len(l.a) {
		na := make([]int, l.sz*2)
		copy(na, l.a)
		l.a = na
	}
	l.a[l.sz] = v
	l.sz++
}

func (l *intList) toArray() []int {
	r := make([]int, l.sz)
	copy(r, l.a[:l.sz])
	return r
}

// mapArr 压缩的有序 (key→count) 数组，二分查询
type mapArr struct {
	keys []int
	cnts []int
}

func newMapArr(keys, cnts []int) *mapArr {
	return &mapArr{keys: keys, cnts: cnts}
}

func (m *mapArr) getCount(key int) int {
	idx := sort.SearchInts(m.keys, key)
	if idx < len(m.keys) && m.keys[idx] == key {
		return m.cnts[idx]
	}
	return 0
}

// --- 块结构 ---

// blocks 一棵满二叉树的块表示
type blocks struct {
	n   int     // 叶子数
	len []int   // B0..Bn 的长度
	val []int   // B0..Bn 的值 (mod p)
}

// buildBlocks 递归生成所有 n 叶满二叉树的块
func buildBlocks(n, p int, pow10 []int, memo [][]*blocks) []*blocks {
	if memo[n] != nil {
		return memo[n]
	}
	var res []*blocks
	if n == 1 {
		res = append(res, &blocks{
			n:   1,
			len: []int{1, 1},
			val: []int{1 % p, 9 % p},
		})
	} else {
		for k := 1; k <= n-1; k++ {
			lefts := buildBlocks(k, p, pow10, memo)
			rights := buildBlocks(n-k, p, pow10, memo)
			for _, L := range lefts {
				for _, R := range rights {
					blen := make([]int, n+1)
					bval := make([]int, n+1)
					// B0 = '1' + L.B0
					concatBlk(1, 1%p, L.len[0], L.val[0], p, pow10, blen, bval, 0)
					// B1..B_{k-1} = L 内部块
					for i := 1; i <= k-1; i++ {
						blen[i] = L.len[i]
						bval[i] = L.val[i]
					}
					// Bk = L.Bk + R.B0
					concatBlk(L.len[k], L.val[k], R.len[0], R.val[0], p, pow10, blen, bval, k)
					// 右内部块
					for j := 1; j <= R.n-1; j++ {
						blen[k+j] = R.len[j]
						bval[k+j] = R.val[j]
					}
					// Bn = R.B_{R.n} + '9'
					concatBlk(R.len[R.n], R.val[R.n], 1, 9%p, p, pow10, blen, bval, n)
					res = append(res, &blocks{n: n, len: blen, val: bval})
				}
			}
		}
	}
	memo[n] = res
	return res
}

func concatBlk(lenA, valA, lenB, valB, p int, pow10 []int, outLen, outVal []int, idx int) {
	outLen[idx] = lenA + lenB
	outVal[idx] = (int)((int64(valA)*int64(pow10[lenB])+int64(valB))%int64(p))
}

// --- DFS 枚举一侧排列 ---

// dfsEnum 枚举从 posStart 开始的 k 个宝石的所有排列，沿途计算块值
func dfsEnum(n, k, posStart int, powL, B, powDi, xi []int, p int,
	t, mask, w int, buckets []*intList, usedMasks []int, usedCnt *int) {
	if t == k {
		lst := buckets[mask]
		if lst == nil {
			lst = newIntList(16)
			buckets[mask] = lst
			usedMasks[*usedCnt] = mask
			*usedCnt++
		}
		lst.add(w)
		return
	}
	pos := posStart + t
	powLj := powL[pos]
	unused := ((1 << n) - 1) ^ mask
	for unused != 0 {
		lb := unused & -unused
		gi := lbToIdx(lb)
		unused ^= lb
		// w2 = ((w * powDi[gi] + xi[gi]) % p * powLj + B[pos]) % p
		w2 := (int(((int64(w)*int64(powDi[gi])+int64(xi[gi]))%int64(p)*
			int64(powLj)+int64(B[pos])) % int64(p)))
		dfsEnum(n, k, posStart, powL, B, powDi, xi, p,
			t+1, mask|(1<<gi), w2, buckets, usedMasks, usedCnt)
	}
}

func lbToIdx(lb int) int {
	idx := 0
	for lb > 1 {
		lb >>= 1
		idx++
	}
	return idx
}

// --- 主函数 ---

func treeOfInfiniteSouls(gem []int, p int, target int) int {
	n := len(gem)
	if n == 0 {
		return 0
	}

	// 预处理：位数、取模值
	d := make([]int, n)    // 每颗宝石的位数
	xi := make([]int, n)   // gem[i] % p
	totalDigits := 0
	for i := 0; i < n; i++ {
		d[i] = digitLen(gem[i])
		xi[i] = gem[i] % p
		if xi[i] < 0 {
			xi[i] += p
		}
		totalDigits += d[i]
	}

	// 10 的幂
	maxExp := totalDigits + 4*n + 10
	pow10 := make([]int, maxExp+1)
	pow10[0] = 1 % p
	for i := 1; i <= maxExp; i++ {
		pow10[i] = int((int64(pow10[i-1]) * 10) % int64(p))
	}
	powDi := make([]int, n)
	for i := 0; i < n; i++ {
		powDi[i] = pow10[d[i]]
	}

	// 每个掩码的位数和
	FULL := (1 << n) - 1
	sumDigMask := make([]int, 1<<n)
	for m := 1; m <= FULL; m++ {
		lb := m & -m
		idx := lbToIdx(lb)
		sumDigMask[m] = sumDigMask[m^lb] + d[idx]
	}

	// 生成所有树形的块
	memo := make([][]*blocks, n+1)
	shapes := buildBlocks(n, p, pow10, memo)

	leftCnt := n / 2
	rightCnt := n - leftCnt
	T := target % p

	var answer int64

	// 逐形状处理
	for _, blk := range shapes {
		B := blk.val
		L := blk.len
		b0 := B[0]

		sumLenLeftBlocks := 0
		for i := 1; i <= leftCnt; i++ {
			sumLenLeftBlocks += L[i]
		}
		sumLenRightBlocks := 0
		for i := leftCnt + 1; i <= n; i++ {
			sumLenRightBlocks += L[i]
		}

		powL := make([]int, n+1)
		for i := 0; i <= n; i++ {
			powL[i] = pow10[L[i]]
		}

		// 右侧：rightCnt 个宝石，块位置 leftCnt+1..n
		rightBuckets := make([]*intList, 1<<n)
		rightUsed := make([]int, 130)
		rightCntVal := 0
		dfsEnum(n, rightCnt, leftCnt+1, powL, B, powDi, xi, p,
			0, 0, 0, rightBuckets, rightUsed, &rightCntVal)

		// 左侧：leftCnt 个宝石，块位置 1..leftCnt
		leftBuckets := make([]*intList, 1<<n)
		leftUsed := make([]int, 130)
		leftCntVal := 0
		dfsEnum(n, leftCnt, 1, powL, B, powDi, xi, p,
			0, 0, 0, leftBuckets, leftUsed, &leftCntVal)

		// 压缩左侧
		leftArr := make([]*mapArr, 1<<n)
		for i := 0; i < leftCntVal; i++ {
			mask := leftUsed[i]
			arr := compressBucket(leftBuckets[mask])
			leftArr[mask] = arr
		}

		// 按右掩码配对补集左掩码
		for i := 0; i < rightCntVal; i++ {
			rMask := rightUsed[i]
			rBucket := rightBuckets[rMask]
			if rBucket == nil {
				continue
			}
			lMask := FULL ^ rMask
			lArr := leftArr[lMask]
			if lArr == nil {
				continue
			}

			sumDigitsRight := sumDigMask[rMask]
			sumDigitsLeft := totalDigits - sumDigitsRight
			M_right := pow10[sumDigitsRight+sumLenRightBlocks]
			M_left := pow10[sumDigitsLeft+sumLenLeftBlocks]
			base := int((int64(int64(b0)*int64(M_left)%int64(p)) * int64(M_right)) % int64(p))

			rArr := compressBucket(rBucket)
			lk := lArr.keys
			lv := lArr.cnts
			for t := 0; t < len(lk); t++ {
				aLeft := lk[t]
				cLeft := lv[t]
				need := (T - base - int(int64(aLeft)*int64(M_right)%int64(p))) % p
				if need < 0 {
					need += p
				}
				cRight := rArr.getCount(need)
				if cRight != 0 {
					answer += int64(cLeft) * int64(cRight)
				}
			}
		}
	}
	return int(answer)
}

// compressBucket 将桶排序压缩为 (keys, cnts)
func compressBucket(lst *intList) *mapArr {
	if lst == nil {
		return nil
	}
	a := lst.toArray()
	sort.Ints(a)
	m := len(a)
	keysTmp := make([]int, m)
	cntsTmp := make([]int, m)
	uniq := 0
	for t := 0; t < m; {
		v := a[t]
		j := t + 1
		for j < m && a[j] == v {
			j++
		}
		keysTmp[uniq] = v
		cntsTmp[uniq] = j - t
		uniq++
		t = j
	}
	return newMapArr(keysTmp[:uniq], cntsTmp[:uniq])
}

func digitLen(x int) int {
	if x == 0 {
		return 1
	}
	cnt := 0
	for x > 0 {
		x /= 10
		cnt++
	}
	return cnt
}
