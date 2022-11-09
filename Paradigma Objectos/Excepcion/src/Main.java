import java.util.List;


public class Main {

    public static void main(String[] args) {

        Desifrador de = new Desifrador();

        String exp = de.solicitarExp();

        List<String> elementos;

        elementos = de.Token(exp);

        int num = de.formular(elementos);

        if (num == -1) {
            System.err.println("El resultado de la operacion es:" + num);
            System.err.println("Operacion invalida.");
        } else {
            System.out.println("El resultado de la operacion es:" + num);
        }
    }

}