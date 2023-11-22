import time


def test():
	s = ''
	for i in range(100000):
		s += 'hello' + str(i)


if __name__ == '__main__':
	start = time.time()
	test()
	end = time.time()
	print(f"(python)耗时={end - start}s")