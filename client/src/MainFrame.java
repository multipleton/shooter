import org.json.simple.parser.ParseException;

import javax.swing.*;
import java.awt.*;
import java.awt.event.*;
import java.io.IOException;
import java.net.URISyntaxException;
import java.util.ArrayList;

public class MainFrame implements KeyListener, MouseListener, MouseMotionListener {

    JFrame f;
    GameField field;
    Runner runner;

    int mouseX;
    int mouseY;

    final int SPEED = 10;

    ArrayList<Integer> pressed;

    MainFrame(Circle player, Circle[] enemies, Bullet[] bullets) throws InterruptedException, IOException, ParseException, URISyntaxException {
        f = new JFrame();
        Lobby lobby = new Lobby(f);
        InitGameBoard(player, enemies, bullets);
    }

    private void InitGameBoard(Circle player, Circle[] enemies, Bullet[] bullets) {
        pressed = new ArrayList<Integer>();

        f.setLayout(null);
        f.setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
        f.addMouseMotionListener(this);
        f.addKeyListener(this);
        f.addMouseListener(this);
        Dimension screenSize = Toolkit.getDefaultToolkit().getScreenSize();

        int screenWidth = screenSize.width;
        int screenHeight = screenSize.height;
        f.setSize(screenWidth, screenHeight);

        field = new GameField(screenWidth, screenHeight);
        field.setPlayer(player);
        if (enemies != null) {
            for (int i = 0; i < enemies.length; i++) {
                field.addEnemy(enemies[i]);
            }
        }

//        field.addBullets(bullets);

        f.add(field);
        f.setVisible(true);
        runner = new Runner(field);
        runner.run();
    }

    @Override
    public void keyPressed(KeyEvent e) {
        int key = e.getKeyCode();
        int dx = 0;
        int dy = 0;

        if (key == KeyEvent.VK_D ||
                key == KeyEvent.VK_A ||
                key == KeyEvent.VK_S ||
                key == KeyEvent.VK_W) {
            pressed.add(key);
        } else return;

        for (int i = 0; i < pressed.size(); i++) {
            if (pressed.get(i) == KeyEvent.VK_D)
                dx = SPEED;
            else if (pressed.get(i) == KeyEvent.VK_A)
                dx = -SPEED;
            else if (pressed.get(i) == KeyEvent.VK_S)
                dy = SPEED;
            else if (pressed.get(i) == KeyEvent.VK_W)
                dy = -SPEED;
        }
        field.updatePlayer(dx, dy, mouseX, mouseY);
    }

    @Override
    public void mouseMoved(MouseEvent e) {
        mouseX = e.getX();
        mouseY = e.getY();
        field.updatePlayer(0, 0, mouseX, mouseY);
    }

    @Override
    public void keyReleased(KeyEvent e) {
        pressed.removeIf(s -> s.equals(e.getKeyCode()));
    }

    @Override
    public void keyTyped(KeyEvent e) {}

    @Override
    public void mouseDragged(MouseEvent e) {}

    @Override
    public void mouseClicked(MouseEvent e) {
//        new Bullet();
    }

    @Override
    public void mousePressed(MouseEvent e) {}

    @Override
    public void mouseReleased(MouseEvent e) {}

    @Override
    public void mouseEntered(MouseEvent e) {}

    @Override
    public void mouseExited(MouseEvent e) {}
}

//        JComboBox<String> servers = new JComboBox<>(new String[] {"server1", "server2", "server3"});
//        JTextField nickname = new JTextField();
//        final JComponent[] inputs = new JComponent[] {
//                new JLabel("Nickname:"),
//                nickname
//        };
//        int result = JOptionPane.showConfirmDialog(null, inputs, "Shooter", JOptionPane.PLAIN_MESSAGE);

//        circle = new DrawCircle(25, 25, screenSize.width, screenSize.height);
//        f.add(circle);
