import matplotlib as mpl
mpl.use('TkAgg')

import matplotlib.pyplot as plt

ns = [2**x for x in range(11)];
data = {
  'goroutines': [6919, 13212, 25469, 50819, 88566, 162391, 299955, 574043, 1129372, 2251411, 4760560],
  'reflection': [10868, 22335, 54882, 148218, 543921, 1694021, 6102920, 22648976, 90204929, 383579039, 1676544681],
  'recursion': [2658, 14707, 44520, 114676, 261880, 560284, 1117642, 2242910, 4784719, 10044186, 20599475],
}
for (label, values) in data.items():
    plt.plot(ns, values, label=label)
plt.legend()
plt.yscale('log')
plt.show()
