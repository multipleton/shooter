public class DrawCircle {

    private int x;
    private int y;
    private final int SIZE = 100;
    private int ring_x = -1;
    private int ring_y = -1;

    public DrawCircle(int initX, int initY) {
        x = initX;
        y = initY;
    }

    public void moveCircle(int new_x, int new_y, int mouseX, int mouseY) {
        x += new_x;
        y += new_y;
        rotateCircle(mouseX, mouseY);
    }

    public void rotateCircle(int new_x, int new_y) {
        double vectorStaticX = 50;
        double vectorStaticY = 0;
        double vectorX = new_x - (x + 50);
        double vectorY = new_y - (y + 50);
        double catet = (vectorStaticX * vectorX) + (vectorStaticY * vectorY);
        double hypo = Math.sqrt(Math.pow(vectorStaticX, 2) + Math.pow(vectorStaticY, 2)) * Math.sqrt(Math.pow(vectorX, 2) + Math.pow(vectorY, 2));
        double cos = catet / hypo;
        double angle = Math.acos(cos);
        double sin;
        if (new_y <= y + 50) {
            sin = Math.sin(-angle);
        }
        else {
            sin = Math.sin(angle);
        }
        ring_x = (int) Math.round((SIZE / 2) * cos) + x + 50;
        ring_y = (int) Math.round((SIZE / 2) * sin) + y + 50;
    }

    public int getX() {
        return this.x;
    }

    public int getY() {
        return this.y;
    }

    public int getSize() {
        return SIZE;
    }

    public int getRingX() {
        return ring_x;
    }

    public int getRingY() {
        return ring_y;
    }
}