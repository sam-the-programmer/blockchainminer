from jax import vmap

def x():
	return 1

g = vmap(x)