/*
    方法一：使用多源的BFS直接广搜即可。
 */
class Solution {
    static class Pos {
        int i;
        int j;
        Pos(int i, int j) {
            this.i = i;
            this.j = j;
        }
        Pos(){
            this.i = -1;
            this.j = -1;
        }
    }
    public int[][] updateMatrix(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        int[][] dist = new int[m][n];
        LinkedList<Pos> queue = new LinkedList<>();

        // init queue and dist
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                if(matrix[i][j] == 0) {
                    queue.add(new Pos(i, j));
                    dist[i][j] = 0;
                }
                else dist[i][j] = -1;
            }
        }

        // BFS
        while (queue.size() > 0) {
            int len = queue.size();
            for(int k=0; k<len; k++) {
                Pos top_node = queue.pop();
                int top_node_dis = dist[top_node.i][top_node.j];
                if(top_node.i>0 && dist[top_node.i-1][top_node.j]==-1) {
                    dist[top_node.i - 1][top_node.j] = top_node_dis + 1;
                    queue.add(new Pos(top_node.i-1, top_node.j));
                }
                if(top_node.i<m-1 && dist[top_node.i+1][top_node.j]==-1) {
                    dist[top_node.i + 1][top_node.j] = top_node_dis + 1;
                    queue.add(new Pos(top_node.i+1, top_node.j));
                }
                if(top_node.j>0 && dist[top_node.i][top_node.j-1]==-1) {
                    dist[top_node.i][top_node.j - 1] = top_node_dis + 1;
                    queue.add(new Pos(top_node.i, top_node.j-1));
                }
                if(top_node.j<n-1 && dist[top_node.i][top_node.j+1]==-1) {
                    dist[top_node.i][top_node.j + 1] = top_node_dis + 1;
                    queue.add(new Pos(top_node.i, top_node.j+1));
                }
            }
        }
        return dist;
    }
}

/*
    方法二： 找到最优子结构，使用动态规划。注意只要用到Integer.MAX_INTEGER就需要考虑是否有越界
 */
class Solution {
    public int[][] updateMatrix(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        int[][] dist = new int[m][n];


        // init and left-top
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                dist[i][j] = matrix[i][j]==0?0:30000;
                if(i>0) dist[i][j] = Math.min(dist[i][j], dist[i-1][j]+1);
                if(j>0) dist[i][j] = Math.min(dist[i][j], dist[i][j-1]+1);
            }
        }

        // right - bottom
        for(int i=m-1; i>=0; i--) {
            for(int j=n-1; j>=0; j--) {
                if(i<m-1) dist[i][j] = Math.min(dist[i][j], dist[i+1][j]+1);
                if(j<n-1) dist[i][j] = Math.min(dist[i][j], dist[i][j+1]+1);
            }
        }

        return dist;
    }
}