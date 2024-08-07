import math
import seaborn as sns

import gym
import torch
import torch.nn as nn
import torch.nn.functional as F

device = torch.device(
    "cuda" if torch.cuda.is_available() else
    "cpu"
)

# TODO: solve CartPole like for (actions)
# - Throttle: [0, 100] percent
# - Aileron
# - Elevator

# Q-learning with shallow neural net as function approximator.
# /delta = Q(s, a) - (r - /gamma max_{a}^{'} Q(s^{'}, a))
# where /delta is temporal difference error and 
# /gamma is the discount factor.

# State Space:
# Speed (Throttle): [0, 1000 or INF]
# Bank Angle (Aileron): [-180, +180]
# Pitch (Elevator): [-90, 90]
# State and Action space are continuous.

env = gym.make("CartPole-v1")

class QN(nn.Module):
    layer_sizes = [32, 64, 32]
    
    # Single variable => num_states = 1 and actions = 1
    def __init__(self, num_states: int, num_actions: int):
        super(QN, self).__init__()

        self.layers = list({
            nn.Linear(num_states, self.layer_sizes[0])
        })

        for i in range(1, len(self.layer_sizes)):
            self.layers.append(nn.Linear(self.layer_sizes[i - 1], self.layer_sizes[i]))

        self.layers.append(nn.Linear(self.layer_sizes[-1], num_actions))

    def forward(self, x):
        for layer in self.layers:
            x = F.relu(layer(x))
        return x

# TODO: add trainer:
# - select or define loss function
# - select or define optimizer
# - create a simulation function or record data from actual flight
# - select action using epsilon greedy approach to feed to sim

# Constants
GAMMA = 1 # discount factor
COMMENCE_EPSILON = 1
DECAY_EPSILON = 1
CULMINATE_EPSILON = 1

LEARNING_RATE = 1

# TODO
