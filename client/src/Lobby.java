import org.json.simple.parser.ParseException;

import javax.swing.*;
import java.io.IOException;
import java.net.URISyntaxException;
import java.util.Arrays;

public class Lobby {

    public Lobby(JFrame frame) throws InterruptedException, URISyntaxException, ParseException, IOException {
        HttpRequestWorker httpRequestWorker = new HttpRequestWorker();

        JTextField nickname = new JTextField();
        final JComponent[] inputs = new JComponent[] {
                new JLabel("Nickname:"),
                nickname
        };
        JOptionPane.showMessageDialog(null, inputs, "Shooter", JOptionPane.PLAIN_MESSAGE);
        if (nickname.getText().length() > 0) {
            long id = httpRequestWorker.Login(nickname.getText());
            Object[] options = {"Join", "Create"};
            int n = JOptionPane.showOptionDialog(frame,
                    "Select how to start game: ",
                    "Shooter",
                    JOptionPane.YES_NO_CANCEL_OPTION,
                    JOptionPane.QUESTION_MESSAGE,
                    null,
                    options,
                    null);
            if (n == 0) {
                String[] titles = httpRequestWorker.GetServers();
                int selectedServer = JOptionPane.showOptionDialog(frame,
                        "Select server: ",
                        "Shooter",
                        JOptionPane.YES_NO_CANCEL_OPTION,
                        JOptionPane.QUESTION_MESSAGE,
                        null,
                        titles,
                        null);
                httpRequestWorker.Join(id, selectedServer + 1);
            }
            else if (n == 1) {
                httpRequestWorker.Create(id);
            }
        }
    }
}
