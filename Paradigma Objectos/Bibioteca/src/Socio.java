import java.util.ArrayList;

public class Socio {

    private int numSocio;
    private String nombre;
    private String direccion;

    ArrayList<Integer> librosEnPrestamo;

    public Socio(int numSocio, String nombre, String direccion) {
        this.numSocio = numSocio;
        this.nombre = nombre;
        this.direccion = direccion;
        this.librosEnPrestamo = new ArrayList<>();
    }


    public int getNumSocio() {
        return numSocio;
    }

    public void setNumSocio(int numSocio) {
        this.numSocio = numSocio;
    }

    public String getNombre() {
        return nombre;
    }


    @Override
    public String toString() {
        return "Socio{" + "numSocio=" + numSocio + ", nombre=" + nombre + ", direccion=" + direccion + '}';
    }


}