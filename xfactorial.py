def factorial(n):
	if n == 1:
		return 1
	else:
		return n*factorial(n-1)

def xfactorial(m):
	if m == 0:
		return 1
	elif m == 1:
		return m
	elif m < 0:
		return None
	else:
		return factorial(m)

print("################ xfactorial(m) ##################")
a = 5
print(xfactorial(a))

	