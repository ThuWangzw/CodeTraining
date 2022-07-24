bool cmp(pair<int, string> a, pair<int, string> b) { 
    if(a.first!=b.first) {
        return a.first>b.first;
    }
    return a.second<b.second;
};

class FoodRatings {
public:
    unordered_map<string, int> foodRate;
    unordered_map<string, string> foodCui;
    unordered_map<string, set<pair<int, string>, decltype(cmp)*>> cuis;
    
    FoodRatings(vector<string> &foods, vector<string> &cuisines, vector<int> &ratings) {
        int n = foods.size();
        for(int i=0; i<n; i++) {
            foodRate[foods[i]] = ratings[i];
            foodCui[foods[i]] = cuisines[i];
            if(cuis.find(cuisines[i]) == cuis.end()) {
                cuis[cuisines[i]] = set<pair<int, string>, decltype(cmp)*>(cmp);
            }
            cuis[cuisines[i]].insert(pair<int, string>(ratings[i], foods[i]));
        }
    }
    
    void changeRating(string food, int newRating) {
        string cui = foodCui[food];
        cuis[cui].erase(cuis[cui].find(pair<int, string>(foodRate[food], food)));
        foodRate[food] = newRating;
        cuis[cui].insert(pair<int, string>(newRating, food));
    }
    
    string highestRated(string cuisine) {
        return (*(cuis[cuisine].begin())).second;
    }
};