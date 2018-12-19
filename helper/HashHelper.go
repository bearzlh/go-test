package helper

type int128 struct {
	i1 int64
	i2 int64
}

func HashCode(channelId string) int64 {
	var b int64 = 0x00000000FFFFFFFF
	var c int64 = 0x0000000080000000
	var d int64 = 0xFFFFFF00
	var f int64 = 0x00000000FFFC0000
	var g int64 = 0x00000000FFFFFFF0
	var j int64 = 0xFFFFFFFF80000000
	var h int64 = 0
	var a int64 = 0
	for i := 0; i < len(channelId); i++ {
		h = 31*h + int64(rune(channelId[i]));
		h = h & b;
	}
	h = ((-1 - h) & b) + ((h << 21) & b); // h = (h << 21) - h - 1;
	h = h & b;
	if (((h & c) >> 31) == 1) {
		a = h>>24 | d;
	} else {
		a = h >> 24;
	}
	a = a & b;
	h = h ^ a;
	h = h & b;
	h = ((h + (h<<3)&b) & b) + ((h << 8) & b); // h * 265
	h = h & b;
	if (((h & c) >> 31) == 1) {
		a = (h >> 14) | f;
	} else {
		a = h >> 14;
	}
	a = a & b;
	h = h ^ a;
	h = h & b;
	h = ((h + ((h << 2) & b)) & b) + ((h << 4) & b); // h * 21
	h = h & b;
	if (((h & c) >> 31) == 1) {
		a = (h >> 28) | g;
	} else {
		a = h >> 28;
	}
	a = a & b;
	h = h ^ a;
	h = h & b;
	h = h + ((h << 31) & b);
	h = h & b;
	if ((h&b)>>31 == 1) {
		return h | j;
	} else {
		return h;
	}
}