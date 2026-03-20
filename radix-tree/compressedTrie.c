#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Edge {
  char *label;
  struct Node *child;
} Edge;

typedef struct Node {
  struct Edge *edges;
  int edge_count;
  int is_terminal;
} Node;

int lcp(char *a, char *b);
Node *new_node();
void add_edge(Node *node, char *label, Node *child);
void insert(Node *node, char *newWord);
void print_tree(Node *node, int depth);

int main() {
  Node *root = new_node();
  insert(root, "cat");
  insert(root, "car");
  insert(root, "cart");
  print_tree(root, 0);
}

int lcp(char *a, char *b) {
  int i = 0;
  while (a[i] != '\0' && b[i] != '\0') {
    if (a[i] != b[i])
      break;
    i++;
  }
  return i;
}

Node *new_node() {
  Node *node = malloc(sizeof(Node));
  node->edges = NULL;
  node->edge_count = 0;
  node->is_terminal = 0;
  return node;
}

void add_edge(Node *node, char *label, Node *child) {
  node->edges = realloc(node->edges, sizeof(Edge) * (node->edge_count + 1));
  node->edges[node->edge_count].label = label;
  node->edges[node->edge_count].child = child;
  node->edge_count++;
}

void print_tree(Node *node, int depth) {
  for (int i = 0; i < node->edge_count; i++) {
    Edge *edge = &node->edges[i];
    for (int j = 0; j < depth; j++) {
      printf("   ");
    }
    if (i == node->edge_count - 1)
      printf("└──"); // last child/ terminal node
    else
      printf("├──"); // other children
    printf("%s", edge->label);
    if (edge->child->is_terminal) {
      printf(" (T)");
    }
    printf("\n");
    print_tree(edge->child, depth + 1);
  }
}

void insert(Node *node, char *newWord) {
  for (int i = 0; i < node->edge_count; i++) {
    Edge *edge = &node->edges[i];
    int k = lcp(edge->label, newWord);
    if (k == 0)
      continue;

    // full match
    if (k == strlen(edge->label)) {
      char *new_suffix = newWord + k;
      // edge case -> "cat" and "cat"
      if (*new_suffix == '\0') {
        edge->child->is_terminal = 1;
        return;
      }
      insert(edge->child, new_suffix);
      return;
    }

    // partial_match
    if (k < strlen(edge->label)) {
      char *common = strndup(edge->label, k);
      char *old_suffix = strdup(edge->label + k);
      char *new_suffix = strdup(newWord + k);

      Node *split = new_node();

      Node *old_child = edge->child;

      add_edge(split, old_suffix, old_child);

      if (*new_suffix == '\0') {
        // edge case: "cat" and "ca"
        split->is_terminal = 1;
        free(new_suffix);
      } else {
        Node *new_child = new_node();
        new_child->is_terminal = 1;

        add_edge(split, new_suffix, new_child);
      }

      free(edge->label);
      edge->label = common;
      edge->child = split;
      return;
    }
  }

  Node *new_child = new_node();
  new_child->is_terminal = 1;
  add_edge(node, strdup(newWord), new_child);
}
