public class DrawCircle {

    public int x;
    public int y;
    public int width = 100;
    public int height = 100;
    public int ring_x = -1;
    public int ring_y = -1;

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
        ring_x = (int) Math.round((width / 2) * cos) + x + 50;
        ring_y = (int) Math.round((width / 2) * sin) + y + 50;
    }
}