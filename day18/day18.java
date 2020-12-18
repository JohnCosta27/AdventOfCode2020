import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class day18 {

    public static void main(String[] args) {
        
        List<int[]> points = new ArrayList<>();

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

        //points = step(points);

        int counter = 0;
        for (int[] p : points) {
            if (p[3] == 1) counter++;
            System.out.printf("x: %d | y: %d | z: %d | State: %d \n", p[0], p[1], p[2], p[3]);
        }
        System.out.println("Part 1: " + counter);

    }
    
    static List<int[]> step(List<int[]> points) {

        List<int[]> newPoints = new ArrayList<>();
        
        for (int[] p : points) {

            int active = 0;

            for (int x = -1; x <= 1; x++) {
                for (int y = -1; y <= 1; y++) {
                    for (int z = -1; z <= 1; z++) {
                        
                        if (!(x == 0 && y == 0 && z == 0)) {
    
                            int[] potentialNeighboor = {p[0] + x, p[1] + y, p[2] + z, 1};
                            //System.out.printf("x: %d | y: %d | z: %d \n", potentialNeighboor[0], potentialNeighboor[1], potentialNeighboor[2]);
                            //Checking if potentialNeighboor is in the list, and then checking if its active.
                            int index = contains(potentialNeighboor, points);
                            if (index != -1) {
                                //System.out.printf("x: %d | y: %d | z: %d \n", potentialNeighboor[0], potentialNeighboor[1], potentialNeighboor[2]);
                                if (points.get(index)[3] == 1) active++;
                            }

                            int active2 = 0;

                            for (int x2 = -1; x <= 1; x++)
                                for (int y2 = -1; y <= 1; y++)
                                    for (int z2 = -1; z <= 1; z++) {
                                        
                                        int[] potentialNeighboor2 = {potentialNeighboor[0] + x2,
                                            potentialNeighboor[1] + y2, potentialNeighboor[2] + z2};
                                        int index2 = contains(potentialNeighboor2, points);
                                        if (index2 != -1) {
                                            if (points.get(index)[3] == 1) active2++;
                                        }

                                    }

                            if (active2 == 3) newPoints.add(potentialNeighboor);
                            
                        }
                        
                    }
                }
            }

            if ((active == 2 || active == 3) && p[3] == 1) newPoints.add(p);
            else if ((active < 2 || active > 3) && p[3] == 1) {
                p[3] = 0;
                newPoints.add(p);
            } else if (p[3] == 0 && active == 3) {
                p[3] = 1;
                newPoints.add(p);
            } else {
                System.out.println("Something else");
            }

        }

        return newPoints;

    }

    /*static int activeNeighboors(int[] p) {
        
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
    }*/
    
    static int contains(int[] p, List<int[]> points) {
        for (int[] search : points) {
            if (search[0] == p[0] && search[1] == p[1] && search[2] == p[2]) {
                return points.indexOf(search);
            }
        }
        return -1;
    }

}