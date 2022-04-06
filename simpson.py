import numpy as np
import timeit
from timeit import Timer



def Simpson(function,lim1,lim2,n):
    # function - підінтегральна функція
    # lim1 lim2 - межі інтегрування
    # n - кількість кроків
    # обчислимо крок інтегрування
    h=(lim2-lim1)/n
    # Підсумовуємо всі значення функції множені на відповідні коефіцієнти
    # та множимо суму на крок
    Int=h/3*(function(lim1)+function(lim2)+
             np.sum(np.asarray([function(lim1+h*i) for i in range(1,n,2)]))*4+
             np.sum(np.asarray([function(lim1+h*i) for i in range(2,n,2)]))*2
           )
    return ('Simpson rule', Int)


xlim1=2
xlim2=4
Ntochok=1000
sim = Simpson(lambda x: np.exp(x)*np.sin((3*x**2)-x)+np.log(x+2), xlim1, xlim2, Ntochok)
print(sim)

#print(timeit.timeit(s), number= 2)


#timeit.timeit(stmt=s, number=100000)

test = """
import numpy as np
def Simpson(function,lim1,lim2,n):
    # function - підінтегральна функція
    # lim1 lim2 - межі інтегрування
    # n - кількість кроків
    # обчислимо крок інтегрування
    h=(lim2-lim1)/n
    # Підсумовуємо всі значення функції множені на відповідні коефіцієнти
    # та множимо суму на крок
    Int=h/3*(function(lim1)+function(lim2)+
             np.sum(np.asarray([function(lim1+h*i) for i in range(1,n,2)]))*4+
             np.sum(np.asarray([function(lim1+h*i) for i in range(2,n,2)]))*2
           )
    return ('Simpson rule', Int)


xlim1=2
xlim2=4
Ntochok=1000
sim = Simpson(lambda x: np.exp(x)*np.sin(3*x**2-x)+np.log(x+2), xlim1, xlim2, Ntochok)
"""
t1 = Timer(stmt=test)
print(t1.timeit(number=10)/10)
#print(timeit.timeit(test), number= 2)
