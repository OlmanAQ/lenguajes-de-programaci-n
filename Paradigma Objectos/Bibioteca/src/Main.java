import java.util.ArrayList;

public class Main {

    public static ArrayList<Libro> listaLibros = new ArrayList<>();
    public static ArrayList<Socio> listaSocios = new ArrayList<>();
    public static ArrayList<Prestamo> listaPrestamos = new ArrayList<>();

    public static void main(String[] args) {

        Libro libro1 = new Libro(1, "El Quijote", "Miguel de Cervantes");
        Libro libro2 = new Libro(2, "El principito", "Antoine de Saint-Exupéry");
        Libro libro3 = new Libro(3, "El señor de los anillos", "J. R. R. Tolkien");


        listaLibros.add(libro1);
        listaLibros.add(libro2);
        listaLibros.add(libro3);


        Socio socio1 = new Socio(0, "Olman", "Kooper");
        Socio socio2 = new Socio(1, "Juan", "Perez");
        Socio socio3 = new Socio(2, "Pedro", "Gomez");

        listaSocios.add(socio1);
        listaSocios.add(socio2);
        listaSocios.add(socio3);


        Prestamo.Funciones fun = new Prestamo.Funciones();
        fun.estadoLibro(listaLibros);

        fun.realizarPrestamo(0, 0, listaLibros, listaSocios, listaPrestamos);

        fun.realizarPrestamo(1, 1, listaLibros, listaSocios, listaPrestamos);
        fun.realizarPrestamo(1, 2, listaLibros, listaSocios, listaPrestamos);


        fun.realizarPrestamo(2, 3, listaLibros, listaSocios, listaPrestamos);
        fun.realizarPrestamo(2, 4, listaLibros, listaSocios, listaPrestamos);
        fun.realizarPrestamo(2, 5, listaLibros, listaSocios, listaPrestamos);



        fun.estadoLibro(listaLibros);

        fun.tamañoPrestamosSocio(listaSocios);
    }

}