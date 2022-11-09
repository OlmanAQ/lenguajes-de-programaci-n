import java.util.ArrayList;


public class Hoja {

    private String nombre;
    private ArrayList<Figura> listaFiguras;
    private ArrayList<Texto> listaTextos;
    private ArrayList<Grupo> listaGrupos;

    public Hoja(String nombre) {
        this.nombre = nombre;
        this.listaTextos = new ArrayList<>();
        this.listaFiguras = new ArrayList<>();
        this.listaGrupos = new ArrayList<>();
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public ArrayList<Figura> getListaFiguras() {
        return listaFiguras;
    }

    public ArrayList<Texto> getListaTextos() {
        return listaTextos;
    }

    public ArrayList<Grupo> getListaGrupos() {
        return listaGrupos;
    }

    public void agregarFigura(Figura figura) {
        listaFiguras.add(figura);
    }

    public void agregarTexto(Texto texto) {
        listaTextos.add(texto);
    }

    public void agregarGrupo(Grupo grupo) {
        listaGrupos.add(grupo);
    }

    @Override
    public String toString() {
        return "Hoja {" + "nombre=" + nombre + ", Figuras=" + listaFiguras + ", Textos=" + listaTextos + ", Grupos=" + listaGrupos + '}';
    }


    public static class Circulo extends Figura {

        public Circulo(String nombre) {
            super(nombre);
            radio = 30;
        }




    }

    public static class Cuadrado extends Figura {

        public Cuadrado(String nombre) {
            super(nombre);
            lados = 4;
            ancho = 4;
            largo = 4;
        }



    }

    public static class Rectangulo extends Figura {

        public Rectangulo(String nombre) {
            super(nombre);
            lados = 4;
            ancho = 4;
            altura = 8;
        }



    }
}
