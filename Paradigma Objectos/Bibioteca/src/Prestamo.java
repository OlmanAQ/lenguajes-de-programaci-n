import java.util.ArrayList;
import java.util.Date;

public class Prestamo {

    private int codigoLibro;
    private int numSocio;
    private Date fecha;

    public Prestamo(int codigoLibro, int numSocio, Date fecha) {
        this.codigoLibro = codigoLibro;
        this.numSocio = numSocio;
        this.fecha = fecha;
    }

    public int getCodigoLibro() {
        return codigoLibro;
    }



    @Override
    public String toString() {
        return "Prestamo{" + "codigoLibro=" + codigoLibro + ", numSocio=" + numSocio + ", fecha=" + fecha + '}';
    }


    public static class Funciones {

        public void tamañoPrestamosSocio(ArrayList<Socio> listaSocios) {
            for (Socio socio : listaSocios) {
                if (socio.librosEnPrestamo.size() >= 3) {
                }
            }
        }

        public void estadoLibro(ArrayList<Libro> listaLibros) {
            for (Libro libro : listaLibros) {
                if (libro.isEstado() == true) {
                    System.out.println("El libro: " + libro.getTitulo() + " se encuentra disponible.");
                } else if (libro.isEstado() == false) {
                    System.out.println("El libro: " + libro.getTitulo() + " se encuentra ocupado.");
                }
            }
        }

        public boolean verificarSiExisteLibro(int codigoLibro, ArrayList<Libro> listaLibros) {
            for (Libro libro : listaLibros) {
                if (libro.getCodigo() == codigoLibro) {
                    return true;
                }
            }
            return false;
        }

        public boolean verificarSiExisteSocio(int numSocio, ArrayList<Socio> listaSocios) {
            for (Socio socio : listaSocios) {
                if (socio.getNumSocio() == numSocio) {
                    return true;
                }
            }
            return false;
        }

        public boolean verificacionlibroDisponible(int codigoLibro, ArrayList<Prestamo> listaPrestamo) {
            boolean verificar = false;
            for (Prestamo prestamo : listaPrestamo) {
                if (prestamo.getCodigoLibro() == codigoLibro) {
                    verificar = true;
                    break;
                }
            }
            if (verificar == true) {
                return verificar;
            } else {
                return verificar;
            }
        }

        public void realizarPrestamo(int numSocio, int codigoLibro, ArrayList<Libro> listaLibros,
                                     ArrayList<Socio> listaSocios, ArrayList<Prestamo> listaPrestamo) {

            boolean socioExiste = false;
            boolean libroExiste = false;
            boolean libroPrestado = false;

            libroExiste = verificarSiExisteLibro(codigoLibro, listaLibros);
            socioExiste = verificarSiExisteSocio(numSocio, listaSocios);
            libroPrestado = verificacionlibroDisponible(codigoLibro, listaPrestamo);

            if (libroExiste == true && socioExiste == true && libroPrestado == false) {
                for (Libro libro : listaLibros) {
                    if (libro.getCodigo() == codigoLibro) {
                        libro.setEstado(false);
                    }
                }
                for (Socio socio : listaSocios) {
                    if (socio.getNumSocio() == numSocio) {
                        socio.librosEnPrestamo.add(codigoLibro);
                    }
                }
                Date date = new Date();
                Prestamo prestamo = new Prestamo(codigoLibro, numSocio, date);
                listaPrestamo.add(prestamo);
            } else {

                if (libroExiste == false) {
                    System.out.println("Codigo es erroneo.");
                }

                if (socioExiste == false) {
                    System.out.println("El socio no existe.");
                }

                if (libroPrestado == true) {
                    System.out.println("El libro no esta.");
                }
            }
        }
    }
}