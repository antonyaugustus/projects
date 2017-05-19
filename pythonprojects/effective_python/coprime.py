
def coprime(a,b):
    for i in range(2,min(a,b)+1):
        if a%i == 0 and b%i == 0:
            print('Coprime')
            break
    else:
        print ('Not Coprime')

def coprime2(a,b):
    is_coprime=True
    for i in range(2,min(a,b)+1):
        if a%i == 0 and b%i == 0:
            is_coprime=True
        else:
            is_coprime=False

        return is_coprime

print(coprime2(4,9))
print(coprime2(3,6))