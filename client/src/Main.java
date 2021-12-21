import org.json.simple.parser.ParseException;

import java.io.IOException;
import java.net.URISyntaxException;

public class Main {
    public static void main(String[] args) throws InterruptedException, URISyntaxException, ParseException, IOException {

        Circle player = new Circle(1920 / 4, 1080 / 10);
//        Circle[] enemies = new Circle[] {new Circle(1920 / 4 + 600, 1080 / 10), new Circle(1920 / 4, 1080 / 10 + 600), new Circle(1920 / 4 + 600, 1080 / 10 + 600)};
//        Bullet[] bullets = new Bullet[] {new Bullet(100, 100, 200, 200)};
        new MainFrame(player, null, null);

    }
}
