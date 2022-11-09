public abstract class Figura {

    String nombre;
    public int radio;
    public int lados;
    public int altura;
    public int largo;
    public int ancho;
    public final double PI = 3.14;

    public Figura(String nombre) {
        this.nombre = nombre;
    }




    @Override
    public String toString() {
        return "Figura {" + "nombre=" + nombre + '}';
    }

}