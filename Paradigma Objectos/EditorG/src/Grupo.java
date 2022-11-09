import java.util.ArrayList;

public class Grupo {

    private String nombre;
    private ArrayList<Figura> listaFiguras;
    private ArrayList<Texto> listaTextos;

    public Grupo(String nombre) {
        this.listaTextos = new ArrayList<>();
        this.listaFiguras = new ArrayList<>();
        this.nombre = nombre;
    }



    public void agregarFigura(Hoja.Circulo figura){
        listaFiguras.add(figura);
    }

    public void agregarTexto(Texto texto){
        listaTextos.add(texto);
    }

    @Override
    public String toString() {
        return "Grupo {" + "nombre=" + nombre + ", listaFiguras=" + listaFiguras + ", listaTextos=" + listaTextos + '}';
    }
}
