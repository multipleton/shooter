import javax.swing.*;
import java.awt.*;

public class GameField extends JPanel {
    int screenWidth;
    int screenHeight;

    DrawCircle player;

    GameField(int screenWidth, int screenHeight) {
        this.screenWidth = screenWidth;
        this.screenHeight = screenHeight;
        setSize(screenWidth, screenHeight);
        setVisible(true);
    }

    public void setPlayer(DrawCircle player) {
        this.player = player;
    }

    public void updatePlayer(int new_x, int new_y, int mouseX, int mouseY) {
        this.player.moveCircle(new_x, new_y, mouseX, mouseY);
        repaint();
    }

    @Override
    public void paint(Graphics g) {
        super.paintComponent(g);
        Graphics2D g2d = (Graphics2D) g;
        g2d.drawRect(screenWidth / 4, screenHeight / 10,
                (int)Math.floor(screenWidth / 1.5), (int)Math.floor(screenHeight / 1.3));
        if (player != null) {
            int size = player.getSize();
            int x = player.getX();
            int y = player.getY();
            g2d.drawOval(x, y, size, size);
            g2d.drawLine(x + size / 2,
                    y + size / 2, player.getRingX(), player.getRingY());
        }
    }
}
