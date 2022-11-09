from distutils.log import error


class numero():
    numero = 0
    numeroS =""
    def __init__(self, nums):
        self.numeroS = nums
        self.numero = int(nums)
    def toString(self):
        return self.numeroS
    def getNumero(self):
        return self.numero
    



def token(text):
    newtext=""
    for l in text:
        if (l in "*+/-"):
            newtext=newtext+" "+l+" "
        else:
            newtext=newtext+l
    newtext=newtext+" "
    return newtext

def listar(text):
    lista=[]
    objeto=""
    for element in text:
        if element == " ":
            try:
                n=numero(objeto)
                lista.append(n)
            except ValueError:
                o=OP(objeto)
                lista.append(o)
            objeto=""
        else:
            objeto=objeto+element
    return lista

def validaCaracter(text):
    newTexto=""
    for l in text:
        if l not in "1234567890+-*/":
            print("El texto contiene un caracter invalido")
        else:
            newTexto=newTexto+l
    return newTexto
    
def validaformato(lista):
    newlista=[]
    turno=0
    for l in lista:
        try:
            l.getNumero()
            if turno ==0:
                newlista.append(l)
                turno=1
            else:
                print("Error de formato")
        except AttributeError:
            if turno ==1:    
                l.getOP()
                newlista.append(l)
                turno=0
            else:
                print("Error de formato")
    return newlista
    

def resultado(lista):
    a=0
    b=0
    s=""
    turno=0
    for element in lista:
        if turno==0:
            a=element.getNumero()
        elif(turno==1):
            s=element.getOP()
        elif(turno==2):
            b=element.getNumero()
            if(s=="+"):
                a=a+b
            elif(s=="-"):
                a=a-b
            elif(s=="*"):
                a=a*b
            elif(s=="/"):
                a=a/b
            turno=0
        turno=turno+1
    return a


class OP():
    OP = ""
    def __init__(self, op):
        self.OP = op
    def getOP(self):
        return self.OP

def main():
    texto=input("Ingrese la operacion: ")
    texto=validaCaracter(texto)
    texto=token(texto)
    lista=listar(texto)
    lista=validaformato(lista)
    print(resultado(lista))

if __name__ == "__main__":
    main()
    