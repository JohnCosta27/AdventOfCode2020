import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class day18 {
    
    static List<int[]> points = new ArrayList<>();

    public static void main(String[] args) {
        
        char[][] input = {{'.', '#', '.'},
        {'.', '.', '#'},
        {'#', '#', '#'}};
        
        //initialise
        for (int i = 0; i < input.length; i++) {
            for (int j = 0; j < input[0].length; j++) {
                int[] newPoint = {i, j, 0, 0};
                if (input[j][i] == '#') newPoint[3] = 1;
                points.add(newPoint);
            }
        }

        for (int[] p : points) {
            System.out.printf("x: %d | y: %d | z: %d \n", p[0], p[1], p[2]);
            System.out.println(activeNeighboors(p));
        }

    }
    
    static int activeNeighboors(int[] p) {
        
        int active = 0;

        for (int x = -1; x <= 1; x++) {
            for (int y = -1; y <= 1; y++) {
                for (int z = -1; z <= 1; z++) {
                    
                    if (!(x == 0 && y == 0 && z == 0)) {

                        int[] potentialNeighboor = {p[0] + x, p[1] + y, p[2] + z};
                        //System.out.printf("x: %d | y: %d | z: %d \n", potentialNeighboor[0], potentialNeighboor[1], potentialNeighboor[2]);
                        int index = contains(potentialNeighboor);
                        if (index != -1) {
                            //System.out.printf("x: %d | y: %d | z: %d \n", potentialNeighboor[0], potentialNeighboor[1], potentialNeighboor[2]);
                            if (points.get(index)[3] == 1) active++;
                        }
                        
                    }
                    
                }
            }
        }
        return active;
    }
    
    static int contains(int[] p) {
        for (int[] search : points) {
            if (search[0] == p[0] && search[1] == p[1] && search[2] == p[2]) {
                return points.indexOf(search);
            }
        }
        return -1;
    }

}