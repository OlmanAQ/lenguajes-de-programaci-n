public class Excepciones extends Exception {

    private int codigoError;

    public Excepciones(int codigoError) {
        super();
        this.codigoError = codigoError;
    }

    @Override
    public String getMessage() {

        String mensaje = "";

        switch (codigoError) {
            case 110:
                mensaje = "Error, caracter invalido.";
                break;
            case 111:
                mensaje = "Error, no tiene el formato de matematica infijo.";
                break;
            default:
                throw new AssertionError();
        }

        return mensaje;

    }
}