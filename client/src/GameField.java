import javax.swing.*;
import java.awt.*;
import java.util.ArrayList;

public class GameField extends JPanel {
    int screenWidth;
    int screenHeight;

    Circle player;

    ArrayList<Circle> enemies;

    GameField(int screenWidth, int screenHeight) {
        this.screenWidth = screenWidth;
        this.screenHeight = screenHeight;
        enemies = new ArrayList<>();
        setSize(screenWidth, screenHeight);
        setVisible(true);
    }

    public void setPlayer(Circle player) {
        if (this.player != null) return;
        this.player = player;
    }

    public void updatePlayer(int new_x, int new_y, int mouseX, int mouseY) {
        this.player.moveCircle(new_x, new_y, mouseX, mouseY);
    }

    public void addEnemy(Circle enemy) {
        this.enemies.add(enemy);
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
        if (enemies != null) {
            for (int i = 0; i < enemies.size(); i++) {
                if (enemies.get(i) == null) continue;
                int x = enemies.get(i).getX();
                int y = enemies.get(i).getY();
                int size = enemies.get(i).getSize();
                g2d.drawOval(x, y, size, size);
                g2d.drawLine(x + size / 2,
                        y + size / 2, enemies.get(i).getRingX(), enemies.get(i).getRingY());
            }
        }
    }
}
