

public class Texto {

    private String texto;
    public Texto(String texto) {
        this.texto = texto;
    }

    public void setTexto(String texto) {
        this.texto = texto;
    }

    @Override
    public String toString() {
        return "Texto{" + "texto=" + texto + '}';
    }

}