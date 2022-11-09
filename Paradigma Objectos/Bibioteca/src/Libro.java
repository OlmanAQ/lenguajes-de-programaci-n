public class Libro {

    private int codigo;
    private String titulo;
    private String autor;
    private boolean estado = true;

    public Libro(int codigo, String titulo, String autor) {
        this.codigo = codigo;
        this.titulo = titulo;
        this.autor = autor;
    }

    public int getCodigo() {
        return codigo;
    }


    public String getTitulo() {
        return titulo;
    }


    public boolean isEstado() {
        return estado;
    }

    public void setEstado(boolean estado) {
        this.estado = estado;
    }

    @Override
    public String toString() {
        return "Libro{" + "codigo=" + codigo + ", titulo=" + titulo + ", autor=" + autor + ", estado=" + estado + '}';
    }
}
