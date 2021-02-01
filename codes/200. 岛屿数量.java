/*
    方法一：比较明显的使用BFS或者DFS来进行搜索
 */
class BFSSolution {
    class Pair {
        int i;
        int j;
        Pair(int i, int j) {
            this.i = i;
            this.j = j;
        }
    }
    public int numIslands(char[][] grid) {
        int cnt = 0;
        int m = grid.length;
        int n = grid[0].length;
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                if(grid[i][j] == '1') {
                    // find an island
                    cnt++;
                    ArrayList<Pair> stack = new ArrayList<>();
                    stack.add(new Pair(i, j));
                    while (stack.size() > 0) {
                        Pair pair = stack.remove(stack.size()-1);
                        grid[pair.i][pair.j] = '2';
                        if(pair.i>0 && grid[pair.i-1][pair.j]=='1') stack.add(new Pair(pair.i-1, pair.j));
                        if(pair.i<m-1 && grid[pair.i+1][pair.j]=='1') stack.add(new Pair(pair.i+1, pair.j));
                        if(pair.j>0 && grid[pair.i][pair.j-1]=='1') stack.add(new Pair(pair.i, pair.j-1));
                        if(pair.j<n-1 && grid[pair.i][pair.j+1]=='1') stack.add(new Pair(pair.i, pair.j+1));
                    }
                }
            }
        }
        return cnt;
    }
}

/*
    方法二：该问题和“动态连通性”相关，因此可以使用并查集
 */
class UnionFindSet {
    private int[] parent;
    private int[] rank;
    UnionFindSet(char[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        parent = new int[m*n];
        rank = new int[m*n];
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                if(grid[i][j] == '1') parent[i*n+j] = i*n+j;
                else parent[i*n+j] = -1;
                rank[i*n+j] = 0;
            }
        }
    }

    Integer find(Integer i) {
        if(!i.equals(parent[i])) {
            parent[i] = find(parent[i]);
            return parent[i];
        }
        return parent[i];
    }

    void union(Integer i, Integer j) {
        Integer root_i = find(i);
        Integer root_j = find(j);
        if(!i.equals(j)) {
            if(rank[root_i] < rank[root_j]) {
                parent[root_i] = root_j;
            }
            else if(rank[root_i] > rank[root_j]) {
                parent[root_j] = root_i;
            }
            else {
                parent[root_i] = root_j;
                rank[root_j]++;
            }
        }
    }

    Integer RootCount() {
        int cnt = 0;
        HashSet<Integer> set= new HashSet<>();
        for (Integer integer : parent) {
            if (integer != -1) {
                Integer root = find(integer);
                if (!set.contains(root)) {
                    set.add(root);
                    cnt++;
                }
            }
        }
        return cnt;
    }
}

class Solution {
    public int numIslands(char[][] grid) {
        UnionFindSet uf = new UnionFindSet(grid);
        int m = grid.length;
        int n = grid[0].length;
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                if(grid[i][j] == '1') {
                    if(i>0 && grid[i-1][j]=='1') uf.union(i*n+j, (i-1)*n+j);
                    if(j>0 && grid[i][j-1]=='1') uf.union(i*n+j, i*n+j-1);
                }
            }
        }
        return uf.RootCount();
    }
}