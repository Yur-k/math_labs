from distutils.command.build_scripts import first_line_re
import numpy as np




def Simpson(function,lim1,lim2,n):
    # function - підінтегральна функція
    # lim1 lim2 - межі інтегрування
    # n - кількість кроків
    # обчислимо крок інтегрування
    h=(lim2-lim1)/n
    print("h ============= ",h)
    # Підсумовуємо всі значення функції множені на відповідні коефіцієнти
    # та множимо суму на крок
    first = function(lim1)+function(lim2)
    print("firstPart ============= ", first)
    pair = [i for i in range(2, n, 2)]
    unpair = [i for i in range(1, n, 2)]
    print("Pair ============= ", pair)
    print("UnPair ============= ", unpair)

    secondV = [function(lim1+h*i) for i in pair]
    thirdV = [function(lim1+h*i) for i in unpair]
    print("pairIndexValue ============= ", secondV)
    print("unpairIndexValue ============= ", thirdV)
    second = np.sum(np.asarray([function(lim1+h*i) for i in range(2,n,2)]))*2
    third = np.sum(np.asarray([function(lim1+h*i) for i in range(1,n,2)]))*4
    print("thirdPart ============= ", third)

    # print("++++++++++++=== ", np.asarray([i for i in range(2,n,2)]))
    # print("Range pair ", np.asarray([function(lim1+h*i) for i in range(2,n,2)]))
    # print("Range unpair ", np.asarray([function(lim1+h*i) for i in range(1,n,2)]))

    print(first,second, third)
    
    Int=h/3*(first + second + third)
    return ('Simpson rule', Int)


xlim1=2
xlim2=4
Ntochok=1000
sim = Simpson(lambda x: np.exp(x)*np.sin((3*x**2)-x)+np.log(x+2), xlim1, xlim2, Ntochok)
print(sim)

#print(timeit.timeit(s), number= 2)


#timeit.timeit(stmt=s, number=100000)

#print(timeit.timeit(test), number= 2)
