public class Main {

    public static void main(String[] args) {

        Hoja h1 = new Hoja("hoja 1");

        Hoja.Cuadrado cu1 = new Hoja.Cuadrado("cuadrado 1");

        h1.agregarFigura(cu1);


        Texto text1 = new Texto("texto 1");
        text1.setTexto(" ABCD ");

        Texto texto2 = new Texto("texto 2");
        texto2.setTexto(" 1234 ");

        h1.agregarTexto(texto2);

        Hoja.Circulo cir1 = new Hoja.Circulo("circulo 1");
        Hoja.Rectangulo rect1 = new Hoja.Rectangulo("rectangulo 1");

        Grupo gr1 = new Grupo("grupo 1");

        gr1.agregarFigura(cir1);
        gr1.agregarTexto(text1);

        h1.agregarGrupo(gr1);
        h1.agregarFigura(rect1);

        System.out.println(h1);


    }

    public abstract static class Figura {

        String nombre;
        public int radio;
        public int lados;
        public int altura;
        public int largo;
        public int ancho;
        public final double PI = 3.14;



    }
}