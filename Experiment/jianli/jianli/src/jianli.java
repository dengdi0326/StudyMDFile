/**
 * Created by Spaxry on 2018/5/25.
 */
import java.awt.*;
import javax.swing.*;
import java.awt.event.*;
import java.io.*;
import java.applet.Applet;
import java.applet.AudioClip;
import java.net.URL;
import java.net.MalformedURLException;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;


class Student {
    String name;
    String sex;
    String birth;
    String major;
    String intro;

}

class Jianli1 extends JFrame{
    static int  i = 0;
    Student stu;
    Student[] student;
    Container container;
    JButton next;
    JPanel p, p1, p2, p3, p4;
    JLabel label, labelName, labelSex, tName, tSex, labelBirth, tBirth, major, tMajor, intro, tIntro, jlpic;

    void setStu(Student a){
        stu = a ;
    }

    Jianli1(String title,Student a){
        super(title);
        stu = a;
        container=this.getContentPane();
        container.setBackground(Color.white);
        container.setLayout(null);
        this.setSize(360,500);

        label=new JLabel("个人简历",JLabel.CENTER);
        label.setFont(new java.awt.Font("Dialog",1,30));//设置个人简历该标签字体的样式为粗体，字号30；

        p=new JPanel(new FlowLayout(FlowLayout.CENTER));
        p.add(label);
        p.setBackground(Color.orange);
        container.add(p);  //设置标签的布局；
        p.setBounds(80,5,200,50);

        labelName = new JLabel("姓名：  ");
        tName = new JLabel(stu.name);

        labelSex = new JLabel("性别:   ");
        tSex = new JLabel(stu.sex);

        p1=new JPanel(new BorderLayout());
        JPanel top = new JPanel(new FlowLayout(FlowLayout.LEFT));
        p1.add(top,BorderLayout.NORTH);
        p1.setBounds(5,70,350,180);
        top.add(labelName);
        top.add(tName);
        top.add(labelSex);
        top.add(tSex);

        labelBirth = new JLabel("生日：   ");
        tBirth = new JLabel(stu.birth);
        top.add(labelBirth);
        top.add(tBirth);

        major = new JLabel("专业：  ");
        tMajor = new JLabel(stu.major);

        JPanel left = new JPanel();
        left.add(major);
        left.add(tMajor);
        p1.add(left, BorderLayout.WEST);

        if(i ==0){
            ImageIcon icon = new ImageIcon("/Users/yusank/Desktop/JAVA/htt1.GIF");
            icon.setImage(icon.getImage().getScaledInstance(icon.getIconWidth()/2,
                    icon.getIconHeight()/2, Image.SCALE_DEFAULT));
            System.out.println(icon.getIconHeight() + "" + icon.getIconWidth());
            jlpic = new JLabel();
            jlpic.setIcon(icon);
            p1.add(jlpic, BorderLayout.EAST);
        }else if(i == 1){
            ImageIcon icon = new ImageIcon("/Users/yusank/Desktop/JAVA/htt2.jpg");
            icon.setImage(icon.getImage().getScaledInstance(icon.getIconWidth()/3,
                    icon.getIconHeight()/3, Image.SCALE_DEFAULT));
            System.out.println(icon.getIconHeight() + "" + icon.getIconWidth());
            jlpic = new JLabel();
            jlpic.setIcon(icon);
            p1.add(jlpic, BorderLayout.EAST);
        }else{
            ImageIcon icon = new ImageIcon("/Users/yusank/Desktop/JAVA/htt3.jpg");
            icon.setImage(icon.getImage().getScaledInstance(icon.getIconWidth(),
                    icon.getIconHeight(), Image.SCALE_DEFAULT));
            System.out.println(icon.getIconHeight() + "" + icon.getIconWidth());
            jlpic = new JLabel();
            jlpic.setIcon(icon);
            p1.add(jlpic, BorderLayout.EAST);
        }
        p2=new JPanel(new FlowLayout(FlowLayout.LEFT));
        p2.setBounds(5,270,350,100);
        intro = new JLabel("自我介绍：     ");
        tIntro = new JLabel(stu.intro);
        p2.add(intro);
        p2.add(tIntro);
        p3=new JPanel(new FlowLayout(FlowLayout.CENTER));
        p3.setBounds(5,350,350,150);
        next = new JButton("下一个");
        p3.add(next);
        next.addActionListener(new React());
        container.add(p1);
        container.add(p2);
        container.add(p3);
        this.setVisible(true);
        this.addWindowListener(new myWindows());
    }
    private class myWindows extends WindowAdapter{
        public void windowClosing(WindowEvent a){
            System.exit(0);
        }
    }
    private class React implements ActionListener{
        public void actionPerformed(ActionEvent e){
            if(i == 0){
                i++;
                try{
                    FileReader fr = new FileReader ("/Users/yusank/Desktop/JAVA/testjianli.txt");
                    BufferedReader br = new BufferedReader(fr);
                    String line = br.readLine();
                    stu.name = line;
                    stu.sex = br.readLine();
                    stu.birth = br.readLine();
                    stu.major = br.readLine();
                    stu.intro = br.readLine();
                    br.close();
                    fr.close();
                    System.out.println("again");
                }
                catch(IOException q){
                    System.out.println("error1");
                }}else{
                i=-1;
                try{
                    FileReader fr = new FileReader ("/Users/yusank/Desktop/JAVA/testjianli.txt");
                    BufferedReader br = new BufferedReader(fr);
                    String line = br.readLine();
                    stu.name = line;
                    stu.sex = br.readLine();
                    stu.birth = br.readLine();
                    stu.major = br.readLine();
                    stu.intro = br.readLine();
                    stu.name = br.readLine();
                    stu.sex = br.readLine();
                    stu.birth = br.readLine();
                    stu.major = br.readLine();
                    stu.intro = br.readLine();
                    br.close();
                    fr.close();
                    System.out.println("again2");
                }
                catch(IOException w){
                    System.out.println("error2");
                }
            }
            new Jianli1("MY ASSUME",stu);
        }
    }

}

public class jianli {

    public static void main (String[] args){

        Student stu1 = new Student();

            Connection con;
            String driver="com.mysql.jdbc.Driver";

            String url="jdbc:mysql://localhost:3306/student";//数据库名字
            String user="";                                  //数据库用户名
            String password="123456";                        //数据库密码
            try {
                Class.forName(driver);
                con = DriverManager.getConnection(url, user, password);
                if (!con.isClosed()) {
                    System.out.println("数据库连接成功");
                }
                Statement statement = con.createStatement();
                String sql = "select * from STU;";
                ResultSet resultSet = statement.executeQuery(sql);
                while (resultSet.next()) {
                    stu1.name = resultSet.getString("name");
                    stu1.sex = resultSet.getNString("sex");
                    stu1.birth = resultSet.getNString("birthday");
                    stu1.intro = resultSet.getNString("intro");
                    stu1.major = resultSet.getNString("major");
                }
                resultSet.close();
                con.close();

            } catch (ClassNotFoundException e) {
                System.out.println("数据库驱动没有安装");

            } catch (SQLException e) {
                System.out.println("数据库连接失败");
            }

        new Jianli1("MY ASSUME",stu1);

        try
        {
            AudioClip music=Applet.newAudioClip(new URL("file:/Users/yusank/Desktop/JAVA/music.wav"));
            music.play();
            System.out.println("soga");
        }
        catch(MalformedURLException e)
        {
            System.err.println(e);
            e.printStackTrace();
        }
    }
}
