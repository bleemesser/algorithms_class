#include <iostream>
#include <vector>
#include <unordered_map>
#include <chrono>

struct Prices
{
    int ore_robot;
    int clay_robot;
    std::pair<int, int> obsidian_robot;
    std::pair<int, int> geode_robot;
};

class State
{
public:
    int turn;
    int ore_robots;
    int clay_robots;
    int obsidian_robots;
    int geode_robots;
    int ore;
    int clay;
    int obsidian;
    int geode;

    State()
    {
        turn = 0;
        ore_robots = 1;
        clay_robots = 0;
        obsidian_robots = 0;
        geode_robots = 0;
        ore = 0;
        clay = 0;
        obsidian = 0;
        geode = 0;
    }

    void step()
    {
        ore += ore_robots;
        clay += clay_robots;
        obsidian += obsidian_robots;
        geode += geode_robots;
        turn++;
    }

    bool can_build_ore_robot(Prices p)
    {
        return ore >= p.ore_robot;
    }

    bool can_build_clay_robot(Prices p)
    {
        return ore >= p.clay_robot;
    }

    bool can_build_obsidian_robot(Prices p)
    {
        return ore >= p.obsidian_robot.first && clay >= p.obsidian_robot.second;
    }

    bool can_build_geode_robot(Prices p)
    {
        return ore >= p.geode_robot.first && obsidian >= p.geode_robot.second;
    }

    void build_ore_robot(Prices p)
    {
        ore -= p.ore_robot;
        ore_robots++;
        step();
    }

    void build_clay_robot(Prices p)
    {
        ore -= p.clay_robot;
        clay_robots++;
        step();
    }

    void build_obsidian_robot(Prices p)
    {
        ore -= p.obsidian_robot.first;
        clay -= p.obsidian_robot.second;
        obsidian_robots++;
        step();
    }

    void build_geode_robot(Prices p)
    {
        ore -= p.geode_robot.first;
        obsidian -= p.geode_robot.second;
        geode_robots++;
        step();
    }

    // create a key for the cache
    int key() {
        return ore_robots * 100000000 + clay_robots * 1000000 + obsidian_robots * 10000 + geode_robots * 100 + ore * 10 + clay * 10 + obsidian * 10 + geode;
    }


};
long cache_hits;
long states;
int solve(State s, Prices p, int turns, std::unordered_map<int, std::pair<int, int>>& cache)
{   
    states++;
    if (s.turn == turns) // if we are at the end of the game, return the score
    {
        cache[s.key()] = std::make_pair(s.geode, s.turn);
        return s.geode;
    }

    if (cache.find(s.key()) != cache.end()) // if we have already solved this state, return the cached value. if we are at a later turn, return 0 because we can't do better than the cached value
    {   
        cache_hits++;
        if (s.turn >= cache[s.key()].second)
        {
            return 0;
        }
    }

    int result = 0;

    if (s.can_build_geode_robot(p))
    {
        State s2 = s;
        s2.build_geode_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }
    if (s.can_build_obsidian_robot(p))
    {
        State s2 = s;
        s2.build_obsidian_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }
    if (s.can_build_clay_robot(p))
    {
        State s2 = s;
        s2.build_clay_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }
    if (s.can_build_ore_robot(p))
    {
        State s2 = s;
        s2.build_ore_robot(p);
        result = std::max(result, solve(s2, p, turns, cache));
    }

    State s2 = s; // if we don't build a robot, we just step
    s2.step();
    result = std::max(result, solve(s2, p, turns, cache));

    cache[s.key()] = std::make_pair(result, s.turn);
    return result;
}

int main()
{
    int turns = 15;

    Prices p;
    p.ore_robot = 4;
    p.clay_robot = 2;
    p.obsidian_robot = std::make_pair(3, 7);
    p.geode_robot = std::make_pair(2, 7);

    State s = State();
    
    // perform tests at different turn counts from 15 to 40 in steps of 5
    for (int i = 15; i <= 40; i += 5)
    {
        turns = i;
        std::unordered_map<int, std::pair<int, int>> cache;
        auto start = std::chrono::high_resolution_clock::now();
        int result = solve(s, p, turns, cache);
        auto end = std::chrono::high_resolution_clock::now();
        auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(end - start);
        std::cout << "Turns: " << turns << " Score: " << result << " Time: " << duration.count() << "ms" << std::endl;
        // std::cout << "Cache hits: " << cache_hits << " States: " << states << std::endl;
        // cache_hits = 0;
        // states = 0;
    }




    return 0;
}