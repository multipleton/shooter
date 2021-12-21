public class Runner extends Thread {

    GameField gameField;

    public Runner(GameField gameField) {
        this.gameField = gameField;
    }

    public void run() {
        while (true) {
            try {
                Thread.sleep(6);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            this.gameField.repaint();
        }
    }
}
