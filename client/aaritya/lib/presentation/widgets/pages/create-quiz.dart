import 'package:flutter/material.dart';

class CreateQuizPage extends StatefulWidget {
  @override
  _CreateQuizPageState createState() => _CreateQuizPageState();
}

class _CreateQuizPageState extends State<CreateQuizPage> {
  final _formKey = GlobalKey<FormState>();
  String _quizTitle = '';
  String _quizDescription = '';
  List<Map<String, dynamic>> _questions = [];

  void _addQuestion() {
    setState(() {
      _questions.add({
        'question': '',
        'options': ['', '', '', ''],
        'correctAnswer': 0,
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Create Quiz'),
      ),
      body: Form(
        key: _formKey,
        child: ListView(
          padding: EdgeInsets.all(16.0),
          children: [
            TextFormField(
              decoration: InputDecoration(labelText: 'Quiz Title'),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please enter a title';
                }
                return null;
              },
              onSaved: (value) => _quizTitle = value!,
            ),
            const SizedBox(height: 16),
            TextFormField(
              decoration: InputDecoration(labelText: 'Quiz Description'),
              maxLines: 3,
              onSaved: (value) => _quizDescription = value!,
            ),
            SizedBox(height: 24),
            const Text(
              'Questions',
            ),
            ..._questions.asMap().entries.map((entry) {
              int idx = entry.key;
              Map<String, dynamic> question = entry.value;
              return Card(
                child: Padding(
                  padding: EdgeInsets.all(8.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      TextFormField(
                        decoration: InputDecoration(labelText: 'Question ${idx + 1}'),
                        onChanged: (value) => question['question'] = value,
                      ),
                      ...List.generate(4, (optionIdx) {
                        return Row(
                          children: [
                            Radio<int>(
                              value: optionIdx,
                              groupValue: question['correctAnswer'],
                              onChanged: (value) {
                                setState(() {
                                  question['correctAnswer'] = value;
                                });
                              },
                            ),
                            Expanded(
                              child: TextFormField(
                                decoration: InputDecoration(labelText: 'Option ${optionIdx + 1}'),
                                onChanged: (value) => question['options'][optionIdx] = value,
                              ),
                            ),
                          ],
                        );
                      }),
                    ],
                  ),
                ),
              );
            }).toList(),
            SizedBox(height: 16),
            ElevatedButton(
              onPressed: _addQuestion,
              child: Text('Add Question'),
            ),
            SizedBox(height: 24),
            ElevatedButton(
              onPressed: () {
                if (_formKey.currentState!.validate()) {
                  _formKey.currentState!.save();
                  // TODO: Save the quiz
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Quiz created successfully')),
                  );
                }
              },
              child: Text('Create Quiz'),
            ),
          ],
        ),
      ),
    );
  }
}
