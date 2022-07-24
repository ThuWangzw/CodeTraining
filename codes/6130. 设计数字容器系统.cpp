class NumberContainers {
public:
    unordered_map<int, int> containers;
    unordered_map<int, set<int>> indexes;
    NumberContainers() {
    }
    
    void change(int index, int number) {
        if(containers.find(index) == containers.end()) {
            containers[index] = number;
            if(indexes.find(number) == indexes.end()) {
                indexes[number] = set<int>();
            }
            indexes[number].insert(index);
        } else {
            int origin = containers[index];
            indexes[origin].erase(indexes[origin].find(index));
            containers[index] = number;
            if(indexes.find(number) == indexes.end()) {
                indexes[number] = set<int>();
            }
            indexes[number].insert(index);
        }
    }
    
    int find(int number) {
        if(indexes.find(number) == indexes.end() || indexes[number].size()==0) {
            return -1;
        }
        return *(indexes[number].begin());
    }
};